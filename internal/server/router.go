package server

import "net/http"

func RegisterRoute() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return mux
}
