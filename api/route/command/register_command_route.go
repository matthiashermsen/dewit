package command

import "net/http"

func RegisterCommandRoute(mux *http.ServeMux, commandPattern string, handler http.HandlerFunc) {
	mux.HandleFunc("/command/"+commandPattern, handler)
}
