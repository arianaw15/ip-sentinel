package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if payload == nil {
		return nil
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		return fmt.Errorf("failed to write JSON response: %w", err)
	}
	return nil
}

func WriteError(w http.ResponseWriter, status int, err error) {
	errorResponse := map[string]string{"error": err.Error()}
	WriteJSON(w, status, errorResponse)
}

var Validate = validator.New()
