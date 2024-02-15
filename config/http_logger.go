package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
		start := time.Now()

		rw := &responseWriterWithStatus{ResponseWriter: w}

		next(rw, r)

		duration := time.Since(start)

		logMsg := fmt.Sprintf("HTTP Request: %s %s Status %d, Duration %v\n", r.Method, r.URL.Path, rw.status, duration)

		logEntry := LogEntry{
			Level:      getLogType(rw.status),
			Message:    logMsg,
			StatusCode: rw.status,
			Method:     r.Method,
			Path:       r.URL.Path,
			Duration:   duration.String(),
			HostName:   os.Getenv("HOSTNAME"),
		}

		logJSON, err := json.Marshal(logEntry)
		if err != nil {
			fmt.Println("Error marshaling log entry:", err)
			return
		}

		if VolLogger == nil {
			log.Print(string(logJSON))
		} else {
			VolLogger.Print(string(logJSON))
		}
	}
}
