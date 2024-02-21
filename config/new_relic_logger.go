package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/mux"
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
		newrelic.ConfigCodeLevelMetricsEnabled(true),
	)

	NRapp = nrApp

	if err != nil {
		fmt.Println("Error creating NewRelic app")
	}
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

// NrHttpContextTransaction Define New Relic middleware for instrumentation
func NrHttpContextTransaction(app *newrelic.Application) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, h := newrelic.WrapHandle(app, r.URL.Path, next)
			h.ServeHTTP(w, r)
		})
	}
}

// NrHttpLogMiddleware Define New Relic middleware for logging
func NrHttpLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriterWithStatus{ResponseWriter: w}

		next.ServeHTTP(rw, r)

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

		if VolLogger != nil {
			VolLogger.Print(string(logJSON))
		} else {
			fmt.Println(string(logJSON))
		}
	})
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

type SqlCommands string

func (f SqlCommands) String() string {
	return string(f)
}

const (
	SELECT SqlCommands = "SELECT"
	INSERT SqlCommands = "INSERT"
)

// StartDatastoreNewRelicSegment starts a New Relic segment for a given query within the provided context
// and returns a function to end the segment. This allows deferred ending of the segment.
func StartDatastoreNewRelicSegment(
	ctx context.Context, query, collection string, operation SqlCommands) (func(), error) {
	txn := newrelic.FromContext(ctx)
	if txn == nil {
		return nil, errors.New("newrelic transaction not found in context")
	}

	segment := newrelic.DatastoreSegment{
		StartTime:          txn.StartSegmentNow(),
		Product:            newrelic.DatastorePostgres,
		Collection:         collection,         // E.g., "public.users"
		Operation:          operation.String(), // E.g., "SELECT"
		ParameterizedQuery: query,
		// TODO: fix circular import for postgres.PostgresConnectionConfig
		Host:         os.Getenv("POSTGRES_HOST"),
		PortPathOrID: os.Getenv("POSTGRES_PORT"),
		DatabaseName: os.Getenv("POSTGRES_DB"),
	}

	// Return a function to end the segment
	return segment.End, nil
}

func NewRelicSegment(ctx context.Context) (func(), error) {
	txn := newrelic.FromContext(ctx)
	if txn == nil {
		return nil, errors.New("there is no transaction in context")
	}

	pc, _, _, _ := runtime.Caller(1) // Caller(1) gives the caller of NewRelicSegment
	longName := runtime.FuncForPC(pc).Name()
	funcName := extractFunctionName(longName)

	segment := txn.StartSegment(funcName)

	return segment.End, nil
}

func extractFunctionName(fullName string) string {
	// Find the last "/" and isolate the substring after it
	slashIndex := strings.LastIndex(fullName, "/")
	shortName := fullName
	if slashIndex != -1 {
		shortName = fullName[slashIndex+1:]
	}

	return shortName
}
