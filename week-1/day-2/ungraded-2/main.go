package main

import (
	"ungraded-2/config"
	"ungraded-2/handler"
)

func main() {
	mux, server := config.Mux()
	db := &handler.ConnectDb{Db: config.ConnectDb()}

	mux.HandleFunc("/villains", db.Villains)
	mux.HandleFunc("/heroes", db.Heroes)


	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}