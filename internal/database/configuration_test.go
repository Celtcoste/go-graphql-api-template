package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNilViper(t *testing.T) {
	require := require.New(t)
	_, err := NewConfiguration()
	require.EqualError(err, "viper is nil")
}

func TestNewConfiguration(t *testing.T) {
	require := require.New(t)
	os.Setenv("DATABASE_ADDRESS", "1.1.1.1")
	os.Setenv("DATABASE_NAME", "api-template")
	os.Setenv("DATABASE_USERNAME", "jean")
	os.Setenv("DATABASE_PASSWORD", "meskin")
	configuration, err := NewConfiguration()
	require.NoError(err)
	require.Equal(
		"jean:meskin@tcp(1.1.1.1:3306)/api-template?parseTime=true",
		configuration.URI(),
	)
}
