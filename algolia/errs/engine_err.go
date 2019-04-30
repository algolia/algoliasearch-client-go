package errs

import (
	"fmt"
)

type AlgoliaErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e AlgoliaErr) Error() string {
	return fmt.Sprintf("Algolia API error [%d] %s", e.Status, e.Message)
}

func IsAlgoliaHTTPErrWithCode(err error, httpCode int) (string, bool) {
	if err == nil {
		return "", false
	}
	e, ok := err.(AlgoliaErr)
	if !ok || e.Status != httpCode {
		return "", false
	}
	return e.Message, true
}
