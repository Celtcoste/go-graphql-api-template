package database

import (
	"fmt"

	"github.com/celtcoste/go-graphql-api-template/internal/util"
)

// Configuration provide information relative to
// MySQL database connection.
type Configuration struct {
	Address  string
	Name     string
	Password string
	Port     int
	Username string
}

// URI is a factory method for generating a database URI.
func (configuration *Configuration) URI() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		configuration.Username,
		configuration.Password,
		configuration.Address,
		configuration.Port,
		configuration.Name)
}

// NewConfiguration is a factory function for creating
// a DatabaseConfiguration instance using a viper sub tree.
func NewConfiguration() (configuration *Configuration, err error) {
	configuration = &Configuration{
		Address:  util.GetEnvStr("DATABASE_ADDRESS"),
		Name:     util.GetEnvStr("DATABASE_NAME"),
		Password: util.GetEnvStr("DATABASE_PASSWORD"),
		Port:     util.GetEnvInt("DATABASE_PORT"),
		Username: util.GetEnvStr("DATABASE_USERNAME"),
	}
	// NOTE: add data validation here if needed.
	return configuration, nil
}
