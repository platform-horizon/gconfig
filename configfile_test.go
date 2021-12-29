package gconfig

import (
	"testing"

	"github.com/spf13/viper"
	"gotest.tools/assert"
)

func TestGetConfigurationFile(t *testing.T) {

	type ConfigurationFileVariables struct {
		LogLevel string
		HTTPPort string
	}

	t.Run("Read correctly configuration file and set to config structure", func(t *testing.T) {
		configPath := "."
		configName := "config.test"
		var config ConfigurationFileVariables

		err := GetConfigFromFile(configName, configPath, &config)

		assert.Equal(t, err, nil, "Unexpected error is not nil")
		assert.Equal(t, config.LogLevel, "info", "Unexpected log level value")
		assert.Equal(t, config.HTTPPort, "3000", "Unexpected http port value")
	})

	t.Run("Return error if config file not found at selected path", func(t *testing.T) {
		wrongPath := "/etc/app"
		fileName := "configuration"
		var config ConfigurationFileVariables

		err := GetConfigFromFile(wrongPath, fileName, &config)
		_, ok := err.(viper.ConfigFileNotFoundError)

		assert.Assert(t, !ok)
		assert.Assert(t, err != nil, "Unexpected error is not nil")
	})

	t.Run("Returns if config file not found with selected name", func(t *testing.T) {
		path := "."
		wrongFileName := "wrongFileName"
		var config ConfigurationFileVariables

		err := GetConfigFromFile(path, wrongFileName, &config)
		_, ok := err.(viper.ConfigFileNotFoundError)

		assert.Assert(t, err != nil, "Unexpected error is nil")
		assert.Assert(t, !ok, "")
	})

	t.Run("Returns if output struct does not contain config variabile", func(t *testing.T) {
		type ConfigEnvVars struct{}

		path := "."
		fileName := "config"

		var wrongConfig ConfigEnvVars

		err := GetConfigFromFile(path, fileName, &wrongConfig)

		assert.Assert(t, err != nil, "Unexpected error is nil")
		assert.Equal(t, wrongConfig, ConfigEnvVars{}, "Unexpected not empty config")
	})
}
