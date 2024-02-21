package config

import (
	"log"

	"github.com/newrelic/go-agent/v3/newrelic"
)

var NRapp *newrelic.Application
var VolLogger *log.Logger
