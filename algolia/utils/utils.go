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

func RetryUntil[T any](
	retry func() (*T, error),
	until func(*T, error) bool,
	maxRetries *int,
	initialDelay *time.Duration,
	maxDelay *time.Duration,
) (*T, error) {
	if maxRetries == nil {
		maxRetries = new(int)
		*maxRetries = 50
	}

	if initialDelay == nil {
		initialDelay = new(time.Duration)
		*initialDelay = 200 * time.Millisecond
	}

	if maxDelay == nil {
		maxDelay = new(time.Duration)
		*maxDelay = 5 * time.Second
	}

	for i := 0; i < *maxRetries; i++ {
		res, err := retry()

		if ok := until(res, err); ok {
			return res, nil
		}

		time.Sleep(*initialDelay)
		*initialDelay *= 2
		if *initialDelay > *maxDelay {
			*initialDelay = *maxDelay
		}
	}

	return nil, &errs.WaitError{}
}
