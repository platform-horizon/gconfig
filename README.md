<div align="center">

# Gconfig

[![Build Status](https://github.com/platform-horizon/gconfig/workflows/build/badge.svg)](https://github.com/platform-horizon/gconfig/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/platform-horizon/gconfig)](https://goreportcard.com/report/github.com/platform-horizon/gconfig)
[![Go Reference](https://pkg.go.dev/badge/github.com/platform-horizon/gconfig.svg)](https://pkg.go.dev/github.com/platform-horizon/gconfig)

</div>

GConfig is a go library to handle configuration file and environment variables.

It uses [viper](https://github.com/spf13/viper) and [koanf](https://github.com/knadh/koanf) to handle env vars and configuration file.

## Install

```ssh
go get -u github.com/platform-horizon/gconfig
```

## Usage

### Read Environment variables

```go
type EnvironmentVariables struct {
    LogLevel string
    HTTPPort string
}

var envVariablesConfig = []gconfig.EnvConfig{
    {
        Key:      "LOG_LEVEL",
        Variable: "LogLevel",
        Required: true
    },
    {
        Key:      "HTTP_PORT",
        Variable: "HTTPPort",
    },
}

var env EnvironmentVariables

if err := gconfig.GetEnvVariables(envVariablesConfig, &env); err != nil {
    panic(err.Error())
}
```

### Read from file

```go
type ConfigurationFileVariables struct {
    LogLevel string
    HTTPPort string
}

var configuration ConfigurationFileVariables

if err := gconfig.GetConfigFromFile("config.test", ".", &configuration); err != nil {
    panic(err.Error())
}
    
```

## Versioning

We use [SemVer](https://semver.org/) for versioning. For the versions available,
see the [tags on this repository](https://github.com/platform-horizon/gconfig/tags).

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE.md](LICENSE.md)
file for details