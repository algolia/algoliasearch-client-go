package errs

import (
	"errors"
	"fmt"
)

var (
	ErrMissingObjectID    = errors.New("objectID is missing or empty")
	ErrMissingKeyID       = errors.New("key ID is not set in Key.Value")
	ErrNoMoreHostToTry    = errors.New("all hosts have been contacted unsuccessfully, it can either be a network error or wrong appID/key credentials were used")
	ErrIndexAlreadyExists = errors.New("destination index already exists, please delete it first as the CopyIndex cannot hold the responsibility of modifying the destination index")
	ErrSameAppID          = errors.New("indices cannot target the same application ID, please use Client.CopyIndex for same-app index copy instead")
)

func ErrJSONDecode(data []byte, t string) error {
	return fmt.Errorf("cannot decode value %s to type %s", string(data), t)
}
