package logging_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/appconfiguration/logging"
)

func TestNewLogLevelInvalidError(t *testing.T) {
	invalidLogLevel := "made-up"

	logLevelInvalidError := logging.NewLogLevelInvalidError(invalidLogLevel)

	assert.NotNil(t, logLevelInvalidError)
}

func TestLogLevelInvalidError(t *testing.T) {
	invalidLogLevel := "made-up"

	logLevelInvalidError := logging.NewLogLevelInvalidError(invalidLogLevel)
	actualErrMsg := logLevelInvalidError.Error()
	expectedErrMsg := fmt.Sprintf("Log level '%s' is invalid.", invalidLogLevel)

	assert.Equal(t, expectedErrMsg, actualErrMsg, fmt.Sprintf("Expected error message to be '%s' but got '%s'", expectedErrMsg, actualErrMsg))
}
