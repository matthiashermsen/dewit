package api_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/appconfiguration/api"
)

func TestPortInvalidError(t *testing.T) {
	portInvalidError := api.PortInvalidError{}
	actualErrMsg := portInvalidError.Error()
	expectedErrMsg := "Port must be in range between 1 and 65535."

	assert.Equal(t, expectedErrMsg, actualErrMsg, fmt.Sprintf("Expected error message to be '%s' but got '%s'", expectedErrMsg, actualErrMsg))
}
