package transport

import (
	"fmt"
	"net/http"
)

type AlgoliaError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e AlgoliaError) Error() string {
	return fmt.Sprintf("Algolia API error [%d] %s", e.Status, e.Message)
}

func IsAlgolia404(err error) bool {
	e, ok := err.(AlgoliaError)
	if !ok {
		return false
	}
	return e.Status == http.StatusNotFound
}
