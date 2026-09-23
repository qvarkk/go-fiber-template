package config

type AppEnv string

const (
	EnvDevelopment AppEnv = "development"
	EnvProduction  AppEnv = "production"
)

type LogLevel string

const (
	LogDebug  LogLevel = "debug"
	LogInfo   LogLevel = "info"
	LogWarn   LogLevel = "warn"
	LogError  LogLevel = "error"
	LogDPanic LogLevel = "dpanic"
	LogPanic  LogLevel = "panic"
	LogFatal  LogLevel = "fatal"
)

func (e AppEnv) IsDevelopment() bool { return e == EnvDevelopment }

func (e AppEnv) IsProduction() bool { return e == EnvProduction }

func AllowedEnvironments() []AppEnv {
	return []AppEnv{EnvDevelopment, EnvProduction}
}

func AllowedLogLevels() []LogLevel {
	return []LogLevel{LogDebug, LogInfo, LogWarn, LogError, LogDPanic, LogPanic, LogFatal}
}
