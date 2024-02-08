package query

import "net/http"

func RegisterQueryRoute(mux *http.ServeMux, queryPattern string, handler http.HandlerFunc) {
	mux.HandleFunc("/query/"+queryPattern, handler)
}
