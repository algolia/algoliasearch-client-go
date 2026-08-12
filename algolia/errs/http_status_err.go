package errs

import "fmt"

// HTTPStatusError is returned by streaming requests when the response has a
// non-2xx status code. It carries the status code and the raw response body
// so that callers can decode the API error payload.
type HTTPStatusError struct {
	statusCode    int
	body          []byte
	correlationID string
}

func NewHTTPStatusError(statusCode int, body []byte) *HTTPStatusError {
	return &HTTPStatusError{
		statusCode: statusCode,
		body:       body,
	}
}

// NewHTTPStatusErrorWithCorrelationID builds the status error carrying the
// Correlation-ID header of the failed response. Quote it when contacting
// Algolia support.
func NewHTTPStatusErrorWithCorrelationID(statusCode int, body []byte, correlationID string) *HTTPStatusError {
	return &HTTPStatusError{
		statusCode:    statusCode,
		body:          body,
		correlationID: correlationID,
	}
}

// StatusCode returns the HTTP status code of the response.
func (e *HTTPStatusError) StatusCode() int {
	return e.statusCode
}

// Body returns the raw response body.
func (e *HTTPStatusError) Body() []byte {
	return e.body
}

// CorrelationID returns the Correlation-ID header of the failed response, or
// the empty string when the response did not carry one.
func (e *HTTPStatusError) CorrelationID() string {
	return e.correlationID
}

func (e *HTTPStatusError) Error() string {
	msg := fmt.Sprintf("HTTP %d: %s", e.statusCode, e.body)

	if e.correlationID != "" {
		return fmt.Sprintf("%s (Correlation-ID: %s)", msg, e.correlationID)
	}

	return msg
}

func (e HTTPStatusError) Is(target error) bool {
	_, ok := target.(*HTTPStatusError)

	return ok
}
