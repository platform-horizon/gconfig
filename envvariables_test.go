package gconfig

import (
	"os"
	"testing"

	"gotest.tools/assert"
)

func testSetEnv(t *testing.T, key string, value string) {
	t.Helper()
	err := os.Setenv(key, value)
	assert.Equal(t, err, nil, "Unexpected error setting env variable")
}

func testUnsetEnv(t *testing.T, key string) {
	t.Helper()
	err := os.Unsetenv(key)
	assert.Equal(t, err, nil, "Unexpected error unsetting env variable")
}

func TestGetEnvVariables(t *testing.T) {

	t.Run("Returns env variables", func(t *testing.T) {
		type TestEnvironmentVariables struct {
			LogLevel string
			HTTPPort string
		}

		var envVariablesConfig = []EnvConfig{
			{
				Key:      "LOG_LEVEL",
				Variable: "LogLevel",
			},
			{
				Key:      "HTTP_PORT",
				Variable: "HTTPPort",
			},
		}

		var env TestEnvironmentVariables

		testSetEnv(t, "LOG_LEVEL", "info")
		defer testUnsetEnv(t, "LOG_LEVEL")

		testSetEnv(t, "HTTP_PORT", "3000")
		defer testUnsetEnv(t, "HTTP_PORT")

		err := GetEnvVariables(envVariablesConfig, &env)

		assert.Equal(t, err, nil, "Unexpected error not nil")
		assert.Equal(t, env.LogLevel, "info", "Unexpected log level value")
		assert.Equal(t, env.HTTPPort, "3000", "Unexpected http port value")
	})

	t.Run("Returns env variables with default", func(t *testing.T) {
		type TestEnvironmentVariables struct {
			LogLevel string
			HTTPPort string
		}

		var envVariablesConfig = []EnvConfig{
			{
				Key:          "LOG_LEVEL",
				Variable:     "LogLevel",
				DefaultValue: "log-level-default-value",
			},
			{
				Key:          "HTTP_PORT",
				Variable:     "HTTPPort",
				DefaultValue: "http-port-default-value",
			},
		}

		var env TestEnvironmentVariables

		err := GetEnvVariables(envVariablesConfig, &env)

		assert.Equal(t, err, nil, "Unexpected error not nil")
		assert.Equal(t, env.LogLevel, "log-level-default-value", "Unexpected log level value")
		assert.Equal(t, env.HTTPPort, "http-port-default-value", "Unexpected http port value")
	})

	t.Run("Returns error if required value not set", func(t *testing.T) {

		type TestEnvironmentVariables struct {
			LogLevel string
			HTTPPort string
		}

		var envVariablesConfig = []EnvConfig{
			{
				Key:      "LOG_LEVEL",
				Variable: "LogLevel",
				Required: true,
			},
			{
				Key:      "HTTP_PORT",
				Variable: "HTTPPort",
				Required: true,
			},
		}

		var env TestEnvironmentVariables

		err := GetEnvVariables(envVariablesConfig, &env)

		assert.Assert(t, err != nil, "Unexpected error nil")
		assert.Equal(t, env.HTTPPort, "")
		assert.Equal(t, env.LogLevel, "")
	})
}
