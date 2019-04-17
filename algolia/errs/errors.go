package errs

import (
	"errors"
	"fmt"
)

var (
	MissingObjectID    = errors.New("missing `objectID` field")
	NoMoreHostToTry    = errors.New("all hosts have been contacted unsuccessfully")
	IndexAlreadyExists = errors.New("destination index already exists, please delete it first as the CopyIndex cannot hold the responsibility of modifying the destination index")
	IteratorEnd        = errors.New("iterator reached the last element")
	SameAppID          = errors.New("indices cannot target the same application ID, please use Client.CopyIndex for same-app index copy instead")
)

func ErrJSONDecode(data []byte, t string) error {
	return fmt.Errorf("cannot decode value %s to type %s", string(data), t)
}
