package main

import (
	"log"
	"ungraded-3/config"
	"ungraded-3/handler"
)

func main() {
	router, server := config.SetupServer()
	app := &handler.App{DB: config.ConnectDb()}

	router.PanicHandler = handler.PanicHandler
	router.GET("/inventories", app.GetInventories)
	router.GET("/inventories/:id", app.GetInventoriesId)
	router.POST("/inventories", app.PostInventory)
	router.PUT("/inventories/:id", app.PutInventory)
	router.DELETE("/inventories/:id", app.DeleteInventory)

	log.Fatal(server.ListenAndServe())
}