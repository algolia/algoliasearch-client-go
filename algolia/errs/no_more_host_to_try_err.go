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
	baseErr := "all hosts have been contacted unsuccessfully, it can either be a server or a network error or wrong appID/key credentials were used. If the error persists, please visit our help center https://alg.li/support-unreachable-hosts or reach out to the Algolia Support team: https://alg.li/support"

	if len(e.intermediateNetworkErrors) > 0 {
		return fmt.Errorf("%s %w", baseErr, errors.Join(e.intermediateNetworkErrors...)).Error()
	}

	return fmt.Sprintf("%s You can use 'ExposeIntermediateNetworkErrors: true' in the config to investigate.", baseErr)
}

func (n NoMoreHostToTryError) Is(target error) bool {
	_, ok := target.(*NoMoreHostToTryError)

	return ok
}
