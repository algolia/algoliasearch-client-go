package errs

import (
	"errors"
	"fmt"
)

var (
	ErrMissingObjectID      = errors.New("objectID is missing or empty")
	ErrMissingKeyID         = errors.New("key ID is not set in Key.Value")
	ErrNoMoreHostToTry      = NewNoMoreHostToTryError()
	ErrIndexAlreadyExists   = errors.New("destination index already exists, please delete it first as the CopyIndex cannot hold the responsibility of modifying the destination index")
	ErrSameAppID            = errors.New("indices cannot target the same application ID, please use Client.CopyIndex for same-app index copy instead")
	ErrObjectNotFound       = errors.New("object not found in with search responses' hits list")
	ErrEmptySecuredAPIKey   = errors.New("secured API key cannot be empty")
	ErrInvalidSecuredAPIKey = errors.New("invalid secured API key, please check that the given key is a secured one")
	ErrValidUntilNotFound   = errors.New("no valid `validUntil` parameter found, please make sure the secured API key has been generated correctly")
)

func ErrJSONDecode(data []byte, t string) error {
	return fmt.Errorf("cannot decode value %s to type %s", string(data), t)
}
