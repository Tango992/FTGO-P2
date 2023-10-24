package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandlerNew(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler new")
	w.Write([]byte("Hello world"))
}

func Middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware1")
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", Middleware1(http.HandlerFunc(HandlerNew)))

	log.Fatal(http.ListenAndServe(":8080", mux))
}