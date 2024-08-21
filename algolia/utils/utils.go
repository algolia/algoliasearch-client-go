//nolint:wrapcheck
package utils

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"
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

// QueryParameterToString convert any query parameters to string.
func QueryParameterToString(obj any) string {
	return strings.ReplaceAll(url.QueryEscape(ParameterToString(obj)), "+", "%20")
}

// ParameterToString convert any parameters to string.
func ParameterToString(obj any) string {
	objKind := reflect.TypeOf(obj).Kind()
	if objKind == reflect.Slice {
		var result []string
		sliceValue := reflect.ValueOf(obj)
		for i := 0; i < sliceValue.Len(); i++ {
			element := sliceValue.Index(i).Interface()
			result = append(result, ParameterToString(element))
		}
		return strings.Join(result, ",")
	}

	if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	if objKind == reflect.Struct {
		if actualObj, ok := obj.(interface{ GetActualInstance() any }); ok {
			return ParameterToString(actualObj.GetActualInstance())
		}
	}

	return fmt.Sprintf("%v", obj)
}

func HasKey[T any](m map[string]T, key string) bool {
	_, ok := m[key]
	return ok
}
