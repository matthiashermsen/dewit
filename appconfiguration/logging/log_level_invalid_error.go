package logging

import "fmt"

type LogLevelInvalidError struct {
	logLevel string
}

func NewLogLevelInvalidError(logLevel string) *LogLevelInvalidError {
	return &LogLevelInvalidError{
		logLevel: logLevel,
	}
}

func (e *LogLevelInvalidError) Error() string {
	return fmt.Sprintf("Log level '%s' is invalid.", e.logLevel)
}
