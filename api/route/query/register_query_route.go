package query

import "net/http"

func RegisterQueryRoute(mux *http.ServeMux, queryPattern string, handler http.HandlerFunc) {
	mux.HandleFunc("GET /query/"+queryPattern, handler)
}
