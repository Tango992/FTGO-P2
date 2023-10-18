package main

import (
	"log"
	"ungraded-3/config"
	"ungraded-3/handler"
)


func main() {
	router, server := config.SetupServer()
	db := &handler.NewInventoryHandler{DB: config.ConnectDb()}

	router.PanicHandler = handler.PanicFunc
	router.GET("/inventories", db.GetInventories)
	router.GET("/inventories/:id", db.GetInventoriesId)
	router.POST("/inventories", db.PostInventory)
	router.PUT("/inventories/:id", db.PutInventory)
	router.DELETE("/inventories/:id", db.DeleteInventory)

	log.Fatal(server.ListenAndServe())
}