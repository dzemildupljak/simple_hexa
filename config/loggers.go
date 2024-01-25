package config

import (
	"log"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/logWriter"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var NRapp *newrelic.Application
var VolLogger *log.Logger
var nrLogWriter logWriter.LogWriter
var nrLogger *log.Logger
