package errs

import (
	"errors"
)

var (
	MissingObjectID = errors.New("missing `objectID` field")
	NoMoreHostToTry = errors.New("all hosts have been contacted unsuccessfully")
)
