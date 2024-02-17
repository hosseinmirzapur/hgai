package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeSchema[T comparable](r *http.Request, schema T) (*T, error) {
	err := json.NewDecoder(r.Body).Decode(schema)
	if err != nil {
		return nil, err
	}

	return &schema, nil

}
