package errs

import (
	"errors"
	"fmt"
)

var ErrNoMoreHostToTry = NewNoMoreHostToTryError()

type NoMoreHostToTryError struct {
	intermediateNetworkErrors []error
	correlationID             string
}

func NewNoMoreHostToTryError(errs ...error) *NoMoreHostToTryError {
	return &NoMoreHostToTryError{
		intermediateNetworkErrors: errs,
	}
}

// NewNoMoreHostToTryErrorWithCorrelationID builds the exhaustion error
// carrying the Correlation-ID of the last attempt whose response had one.
// Quote it when contacting Algolia support.
func NewNoMoreHostToTryErrorWithCorrelationID(correlationID string, errs ...error) *NoMoreHostToTryError {
	return &NoMoreHostToTryError{
		intermediateNetworkErrors: errs,
		correlationID:             correlationID,
	}
}

func (e *NoMoreHostToTryError) IntermediateNetworkErrors() []error {
	return e.intermediateNetworkErrors
}

// CorrelationID returns the Correlation-ID header of the last retry attempt whose
// response carried one, or empty when no attempt did.
func (e *NoMoreHostToTryError) CorrelationID() string {
	return e.correlationID
}

func (e *NoMoreHostToTryError) Error() string {
	baseErr := "all hosts have been contacted unsuccessfully, it can either be a server or a network error or wrong appID/key credentials were used. If the error persists, please visit our help center https://alg.li/support-unreachable-hosts or reach out to the Algolia Support team: https://alg.li/support"

	var msg string
	if len(e.intermediateNetworkErrors) > 0 {
		msg = fmt.Errorf("%s %w", baseErr, errors.Join(e.intermediateNetworkErrors...)).Error()
	} else {
		msg = fmt.Sprintf("%s You can use 'ExposeIntermediateNetworkErrors: true' in the config to investigate.", baseErr)
	}

	if e.correlationID != "" {
		return fmt.Sprintf("%s (Correlation-ID: %s)", msg, e.correlationID)
	}

	return msg
}

func (n NoMoreHostToTryError) Is(target error) bool {
	_, ok := target.(*NoMoreHostToTryError)

	return ok
}
