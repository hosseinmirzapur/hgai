package utils

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, data any) {
	jData, err := json.Marshal(data)
	if err != nil {
		HandleError(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jData)
}

func ErrorResponse(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	data := make(map[string]any)
	data["msg"] = msg
	data["code"] = code

	jData, err := json.Marshal(data)
	if err != nil {
		HandleError(err)
		return
	}

	w.Write(jData)
}
