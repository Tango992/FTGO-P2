package main

import (
	"ungraded-2/config"
	"ungraded-2/handler"
)

func main() {
	mux, server := config.Mux()
	app := &handler.App{DB: config.ConnectDb()}

	mux.HandleFunc("/villains", app.Villains)
	mux.HandleFunc("/heroes", app.Heroes)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
