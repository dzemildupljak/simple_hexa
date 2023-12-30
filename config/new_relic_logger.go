package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/logWriter"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var NRapp *newrelic.Application
var nrLogWriter logWriter.LogWriter
var nrLogger *log.Logger

func NewNRApplication() {
	licence, valid := os.LookupEnv("NEW_RELIC_LICENCE")
	if !valid {
		fmt.Println("Invalid licence key")
	}

	NRapp, err := newrelic.NewApplication(
		newrelic.ConfigAppName("myhexaapp"),
		newrelic.ConfigLicense(licence),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		fmt.Println("Error creating NewRelic app")
		return
	}
	nrLogWriter = logWriter.New(os.Stdout, NRapp)
	nrLogger = log.New(&nrLogWriter, "", log.Default().Flags())
}

func NrHttpLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log information about the incoming request.
		start := time.Now()
		// Create a custom ResponseWriter to capture the status code.
		rw := &responseWriterWithStatus{ResponseWriter: w}

		// Call the next handler in the chain.
		next(rw, r)

		// Log information about the outgoing response into NewRelic and stdout.
		duration := time.Since(start)
		nrLogger.Printf("HTTP Request: %s %s Status %d, Duration %v\n", r.Method, r.URL.Path, rw.status, duration)
	}
}
