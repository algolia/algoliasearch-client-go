package errs

import (
	"errors"
	"fmt"
)

var (
	ErrMissingObjectID = errors.New("missing `objectID` field")
	ErrNoMoreHostToTry = errors.New("all hosts have been contacted unsuccessfully")
)

func ErrJSONDecode(data []byte, t string) error {
	return fmt.Errorf("cannot decode value %s to type %s", string(data), t)
}
