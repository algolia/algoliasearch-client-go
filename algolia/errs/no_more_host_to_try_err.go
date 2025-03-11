package errs

import (
	"errors"
	"fmt"
)

var ErrNoMoreHostToTry = NewNoMoreHostToTryError()

type NoMoreHostToTryError struct {
	intermediateNetworkErrors []error
}

func NewNoMoreHostToTryError(errs ...error) *NoMoreHostToTryError {
	return &NoMoreHostToTryError{
		intermediateNetworkErrors: errs,
	}
}

func (e *NoMoreHostToTryError) IntermediateNetworkErrors() []error {
	return e.intermediateNetworkErrors
}

func (e *NoMoreHostToTryError) Error() string {
	if len(e.intermediateNetworkErrors) > 0 {
		return fmt.Errorf("all hosts have been contacted unsuccessfully, it can either be a server or a network error or wrong appID/key credentials were used. %w", errors.Join(e.intermediateNetworkErrors...)).Error()
	}
	return "all hosts have been contacted unsuccessfully, it can either be a server or a network error or wrong appID/key credentials were used. You can use 'ExposeIntermediateNetworkErrors: true' in the config to investigate."
}

func (n NoMoreHostToTryError) Is(target error) bool {
	_, ok := target.(*NoMoreHostToTryError)

	return ok
}
