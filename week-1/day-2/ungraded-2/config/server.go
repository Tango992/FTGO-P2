package config

import "net/http"

func Mux() (*http.ServeMux, *http.Server) {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	return mux, &server
}
