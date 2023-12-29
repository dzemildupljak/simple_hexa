package config

import (
	"fmt"
	"net/http"
	"time"
)

type responseWriterWithStatus struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriterWithStatus) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func HttpLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log information about the incoming request.
		start := time.Now()

		// NrLogWriter.Write("Incoming HTTP Request: %s %s\n", r.Method, r.URL.Path)
		fmt.Printf("Incoming HTTP Request: %s %s\n", r.Method, r.URL.Path)

		// Create a custom ResponseWriter to capture the status code.
		rw := &responseWriterWithStatus{ResponseWriter: w}

		// Call the next handler in the chain.
		next(rw, r)

		// Log information about the outgoing response.
		duration := time.Since(start)

		fmt.Printf("  -Outgoing HTTP Response: Status %d, Duration %v\n", rw.status, duration)
	}
}
