package errs

import "fmt"

// HTTPStatusError is returned by streaming requests when the response has a
// non-2xx status code. It carries the status code and the raw response body
// so that callers can decode the API error payload.
type HTTPStatusError struct {
	statusCode int
	body       []byte
}

func NewHTTPStatusError(statusCode int, body []byte) *HTTPStatusError {
	return &HTTPStatusError{
		statusCode: statusCode,
		body:       body,
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

func (e *HTTPStatusError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.statusCode, e.body)
}

func (e HTTPStatusError) Is(target error) bool {
	_, ok := target.(*HTTPStatusError)

	return ok
}
