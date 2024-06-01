package ping_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/ping"
)

func TestHandle(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/ping", nil)

	assert.NoError(t, err)

	handler := ping.Handle()

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code, fmt.Sprintf("Expected status code '%d', but got '%d'", http.StatusOK, recorder.Code))
}
