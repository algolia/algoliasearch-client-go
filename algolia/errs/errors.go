package errs

import (
	"errors"
)

var (
	ErrMissingObjectID = errors.New("missing `objectID` field")
	ErrNoMoreHostToTry = errors.New("all hosts have been contacted unsuccessfully")
)
