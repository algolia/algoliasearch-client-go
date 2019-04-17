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

func IsAlgoliaHTTPErr(err error, httpCode int) (bool, string) {
	if err == nil {
		return false, ""
	}
	e, ok := err.(AlgoliaErr)
	if !ok || e.Status != httpCode {
		return false, ""
	}
	return true, e.Message
}
