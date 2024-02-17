package utils

import (
	"errors"
	"net/http"
)

func Method(r *http.Request, allowed string) error {
	if r.Method != allowed {
		return errors.New("method not allowed")
	}
	return nil
}
