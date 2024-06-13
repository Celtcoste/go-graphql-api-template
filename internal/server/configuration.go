package server

import (
	"fmt"
	"time"

	"github.com/celtcoste/go-graphql-api-template/internal/util"
)

// Configuration provide options for HTTP
// server exposure.
type Configuration struct {
	Hostname        string
	Playground      bool
	Port            int
	HealthPort      string
	ShutdownTimeout time.Duration
}

// NewConfiguration is a factory function for creating a
// Configuration instance using a viper sub tree.
func NewConfiguration() (configuration *Configuration, err error) {
	configuration = &Configuration{
		Hostname:   util.GetEnvStr("SERVER_HOST"),
		Playground: util.GetEnvBool("SERVER_PLAYGROUND"),
		Port:       util.GetEnvInt("SERVER_PORT"),
		HealthPort: util.GetEnvStr("SERVER_HEALTH_PORT"),
		// TODO: set as configuration value.
		ShutdownTimeout: time.Second * 20,
	}
	// NOTE: add data validation here if needed.
	return configuration, nil
}

// Addr is a factory method for generating a server binding address.
func (configuration *Configuration) Addr() string {
	return fmt.Sprintf(
		"%s:%d",
		configuration.Hostname,
		configuration.Port)
}
