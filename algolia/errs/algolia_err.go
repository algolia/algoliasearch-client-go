package errs

import (
	"fmt"
)

// AlgoliaErr represents error responses coming from the Algolia Search API when
// an HTTP call has been sent but the request could not have been processed by
// the engine.
//
// This error provides more details regarding potential issues with payloads
// being sent to the API or limits and quotas being reached.
type AlgoliaErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e AlgoliaErr) Error() string {
	return fmt.Sprintf("Algolia API error [%d] %s", e.Status, e.Message)
}

// IsAlgoliaErr type-checks the given error against an AlgoliaErr. If the error
// can be converted, the inner AlgoliaErr is returned along with true.
// Otherwise, the returned boolean is false and nil is returned.
func IsAlgoliaErr(err error) (*AlgoliaErr, bool) {
	if err == nil {
		return nil, false
	}

	e, ok := err.(AlgoliaErr)
	eptr, okPtr := err.(*AlgoliaErr)
	if !ok && !okPtr {
		return nil, false
	}

	if okPtr {
		return eptr, true
	}

	return &e, true
}

// IsAlgoliaErrWithCode returns the error message along with true if the given
// error is an AlgoliaErr and its status code matches the given HTTP code. In
// any other case, an empty string and false are returned.
func IsAlgoliaErrWithCode(err error, httpCode int) (string, bool) {
	if e, ok := IsAlgoliaErr(err); ok && e.Status == httpCode {
		return e.Message, true
	}
	return "", false
}
