package route_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/route"
)

func TestChainHandlers(t *testing.T) {
	t.Run("No middlewares.", func(t *testing.T) {
		headerKey := "foo"
		expectedHeaderValue := "bar"

		routeHandler := route.ChainHandlers(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(headerKey, expectedHeaderValue)

			w.WriteHeader(http.StatusOK)
		})

		request, err := http.NewRequest("GET", "/", nil)

		assert.NoError(t, err, "Expected no request error")

		recorder := httptest.NewRecorder()
		routeHandler.ServeHTTP(recorder, request)

		actualHeaderValue := recorder.Header().Get(headerKey)
		assert.Equal(t, expectedHeaderValue, actualHeaderValue, fmt.Sprintf("Expected header value to be '%s', but got '%s'", expectedHeaderValue, actualHeaderValue))
	})

	t.Run("With middlewares in correct execution order.", func(t *testing.T) {
		headerKey := "order"

		routeHandler := func(w http.ResponseWriter, r *http.Request) {
			actualOrder := w.Header().Get(headerKey)
			expectedOrder := "1"

			assert.Equal(t, expectedOrder, actualOrder, fmt.Sprintf("Expected order to be '%s' but got '%s'", expectedOrder, actualOrder))

			w.WriteHeader(http.StatusOK)
		}

		firstMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				w.Header().Add(headerKey, "0")

				next.ServeHTTP(w, r)
			}
		}

		secondMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				actualOrder := w.Header().Get(headerKey)
				expectedOrder := "0"

				assert.Equal(t, expectedOrder, actualOrder, fmt.Sprintf("Expected order to be '%s' but got '%s'", expectedOrder, actualOrder))

				w.Header().Set(headerKey, "1")

				next.ServeHTTP(w, r)
			}
		}

		finalRouteHandler := route.ChainHandlers(routeHandler, firstMiddleware, secondMiddleware)

		request, err := http.NewRequest("GET", "/", nil)

		assert.NoError(t, err, "Expected no request error")

		recorder := httptest.NewRecorder()
		finalRouteHandler.ServeHTTP(recorder, request)
	})
}
