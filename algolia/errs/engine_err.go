package errs

import (
	"fmt"
	"net/http"
)

type EngineErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e EngineErr) Error() string {
	return fmt.Sprintf("Algolia API error [%d] %s", e.Status, e.Message)
}

func IsAlgolia404(err error) bool {
	e, ok := err.(EngineErr)
	if !ok {
		return false
	}
	return e.Status == http.StatusNotFound
}
