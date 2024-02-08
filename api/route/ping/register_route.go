package ping

import "net/http"

func RegisterRoute(mux *http.ServeMux) {
	mux.HandleFunc("/ping", Handle())
}
