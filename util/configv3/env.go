package configv3

import (
	"strconv"
	"strings"
	"time"
)

// EnvOverride represents all the environment variables read by the CF CLI
type EnvOverride struct {
	BinaryName       string
	CFColor          string
	CFDialTimeout    string
	CFHome           string
	CFLogLevel       string
	CFPluginHome     string
	CFStagingTimeout string
	CFStartupTimeout string
	CFTrace          string
	DockerPassword   string
	Experimental     string
	ForceTTY         string
	HTTPSProxy       string
	Lang             string
	LCAll            string
}

// BinaryName returns the running name of the CF CLI
func (config *Config) BinaryName() string {
	return config.ENV.BinaryName
}

// DialTimeout returns the timeout to use when dialing. This is based off of:
//   1. The $CF_DIAL_TIMEOUT environment variable if set
//   2. Defaults to 5 seconds
func (config *Config) DialTimeout() time.Duration {
	if config.ENV.CFDialTimeout != "" {
		envVal, err := strconv.ParseInt(config.ENV.CFDialTimeout, 10, 64)
		if err == nil {
			return time.Duration(envVal) * time.Second
		}
	}

	return DefaultDialTimeout
}

// DockerPassword returns the docker password from the environment.
func (config *Config) DockerPassword() string {
	return config.ENV.DockerPassword
}

// Experimental returns whether or not to run experimental CLI commands. This
// is based off of:
//   1. The $CF_CLI_EXPERIMENTAL environment variable if set
//   2. Defaults to false
func (config *Config) Experimental() bool {
	if config.ENV.Experimental != "" {
		envVal, err := strconv.ParseBool(config.ENV.Experimental)
		if err == nil {
			return envVal
		}
	}

	return false
}

// HTTPSProxy returns the proxy url that the CLI should use. The url is based
// off of:
//   1. The $https_proxy environment variable if set
//   2. Defaults to the empty string
func (config *Config) HTTPSProxy() string {
	if config.ENV.HTTPSProxy != "" {
		return config.ENV.HTTPSProxy
	}

	return ""
}

// LogLevel returns the global log level. The levels follow Logrus's log level
// scheme. This value is based off of:
//   - The $CF_LOG_LEVEL and an int/warn/info/etc...
//   - Defaults to PANIC/0 (ie no logging)
func (config *Config) LogLevel() int {
	if config.ENV.CFLogLevel != "" {
		envVal, err := strconv.ParseInt(config.ENV.CFLogLevel, 10, 32)
		if err == nil {
			return int(envVal)
		}

		switch strings.ToLower(config.ENV.CFLogLevel) {
		case "fatal":
			return 1
		case "error":
			return 2
		case "warn":
			return 3
		case "info":
			return 4
		case "debug":
			return 5
		}
	}

	return 0
}

// StagingTimeout returns the max time an application staging should take. The
// time is based off of:
//   1. The $CF_STAGING_TIMEOUT environment variable if set
//   2. Defaults to the DefaultStagingTimeout
func (config *Config) StagingTimeout() time.Duration {
	if config.ENV.CFStagingTimeout != "" {
		val, err := strconv.ParseInt(config.ENV.CFStagingTimeout, 10, 64)
		if err == nil {
			return time.Duration(val) * time.Minute
		}
	}

	return DefaultStagingTimeout
}

// StartupTimeout returns the max time an application should take to start. The
// time is based off of:
//   1. The $CF_STARTUP_TIMEOUT environment variable if set
//   2. Defaults to the DefaultStartupTimeout
func (config *Config) StartupTimeout() time.Duration {
	if config.ENV.CFStartupTimeout != "" {
		val, err := strconv.ParseInt(config.ENV.CFStartupTimeout, 10, 64)
		if err == nil {
			return time.Duration(val) * time.Minute
		}
	}

	return DefaultStartupTimeout
}
