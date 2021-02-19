package logging

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	LogLevelEnv = "LOG_LEVEL"
	LogColorEnv = "LOG_COLOR"
)

func FromEnv() zerolog.Logger {
	// configure the global logger
	return log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		NoColor:    !Colored(),
	}).Level(Level())
}

func Colored() bool {
	switch strings.ToLower(envDefault(LogColorEnv, "false")) {
	case "true", "t", "yes", "y":
		return true
	}
	return false
}

func Level() zerolog.Level {
	switch strings.ToLower(envDefault(LogLevelEnv, "info")) {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}

func envDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
