package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonHandler struct {
}

func NewHandler() *JsonHandler {
	return &JsonHandler{}
}

func (j *JsonHandler) SetupHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (j *JsonHandler) ValidateBody(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request Body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func (j *JsonHandler) WriteJson(w http.ResponseWriter, status int, v any) error {

	j.SetupHeader(w)
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)

}

func (j *JsonHandler) WriteError(w http.ResponseWriter, status int, v any) {
	j.WriteJson(w, status, v)
}
