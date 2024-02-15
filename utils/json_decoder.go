package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JsonDecode[T any](r *http.Request) (T, error) {
	var v T
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // Call this before Decode
	err := dec.Decode(&v)       // Pass a pointer to v
	if err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}

	return v, nil
}
