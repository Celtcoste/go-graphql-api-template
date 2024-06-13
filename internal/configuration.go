package internal

import (
	"github.com/celtcoste/go-graphql-api-template/internal/database"
	"github.com/celtcoste/go-graphql-api-template/internal/server"
)

// Configuration is the root level configuration holder.
type Configuration struct {
	Database *database.Configuration
	Server   *server.Configuration
}

// NewConfiguration is a factory function for creating a root level
// application Configuration. Evaluating from both YAML file and
// environment variables.
func NewConfiguration() (configuration *Configuration, err error) {
	databaseConfiguration, err := database.NewConfiguration()
	if err != nil {
		return nil, err
	}
	serverConfiguration, err := server.NewConfiguration()
	if err != nil {
		return nil, err
	}
	return &Configuration{
		Database: databaseConfiguration,
		Server:   serverConfiguration,
	}, nil
}
