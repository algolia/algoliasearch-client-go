//nolint:wrapcheck
package utils

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/errs"
)

// ToPtr is a helper routine that returns a pointer to the given value.
func ToPtr[T any](v T) *T {
	return &v
}

type Nullable[T any] struct {
	value *T
	isSet bool
}

func (v Nullable[T]) Get() *T {
	return v.value
}

func (v *Nullable[T]) Set(val *T) {
	v.value = val
	v.isSet = true
}

func (v Nullable[T]) IsSet() bool {
	return v.isSet
}

func (v *Nullable[T]) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullable[T any](val *T) *Nullable[T] {
	return &Nullable[T]{value: val, isSet: true}
}

func (v Nullable[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *Nullable[T]) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// IsNilOrEmpty checks if an input is nil or empty.
func IsNilOrEmpty(i any) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	case reflect.Bool:
		return false
	default:
		return reflect.ValueOf(i).IsZero()
	}
}

type IterableError[T any] struct {
	Validate func(*T, error) bool
	Message  func(*T, error) string
}

func CreateIterable[T any](
	execute func(*T, error) (*T, error),
	validate func(*T, error) bool,
	aggregator func(*T, error),
	timeout func() time.Duration,
	iterableErr *IterableError[T],
) (*T, error) {
	var executor func(*T, error) (*T, error)

	executor = func(previousResponse *T, previousError error) (*T, error) {
		response, responseErr := execute(previousResponse, previousError)

		if aggregator != nil {
			aggregator(response, responseErr)
		}

		if validate(response, responseErr) {
			return response, responseErr
		}

		if iterableErr != nil && iterableErr.Validate(response, responseErr) {
			if iterableErr.Message != nil {
				return nil, errs.NewWaitError(iterableErr.Message(response, responseErr))
			}

			return nil, errs.NewWaitError("an error occurred")
		}

		if timeout == nil {
			timeout = func() time.Duration {
				return 1 * time.Second
			}
		}

		time.Sleep(timeout())

		return executor(response, responseErr)
	}

	return executor(nil, nil)
}
