package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func Json(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	user := User{
		Name: "John",
		Age: 40,
		Email: "john@mail.com",
	}
	
	json.NewEncoder(w).Encode(user)
}

func Html(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1> Hello World</h1>")
}


func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/json", Json)
	mux.HandleFunc("/html", Hello)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
