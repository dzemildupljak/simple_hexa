package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/logWriter"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewNRApplication() {
	licence, valid := os.LookupEnv("NEW_RELIC_LICENCE")
	if !valid {
		fmt.Println("Invalid licence key")
	}

	nrApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName("myhexaapp"),
		newrelic.ConfigLicense(licence),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	NRapp = nrApp

	if err != nil {
		fmt.Println("Error creating NewRelic app")
	}
}

func NewNrLogger() {
	nrLogWriter = logWriter.New(os.Stdout, NRapp)
	nrLogger = log.New(&nrLogWriter, "", 0)
}

type LogEntry struct {
	Level      string `json:"level"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Method     string `json:"method"`
	Path       string `json:"path"`
	Duration   string `json:"duration"`
	HostName   string `json:"hostname"`
}

func NrHttpLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriterWithStatus{ResponseWriter: w}

		next(rw, r)

		duration := time.Since(start)
		log_msg := fmt.Sprintf("HTTP Request: %s %s Status %d, Duration %v\n", r.Method, r.URL.Path, rw.status, duration)
		logEntry := LogEntry{
			Level:      getLogType(rw.status),
			Message:    log_msg,
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

		VolLogger.Print(string(logJSON))
	}
}

func getLogType(code int) string {
	var msg string
	switch {
	case code >= 400:
		msg = "ERROR"
	case code >= 300:
		msg = "WARNING"
	case code >= 200:
		msg = "INFO"
	default:
		msg = "INFO"
	}

	return msg
}
