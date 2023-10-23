package main

import (
	"preview-week2/config"
	"preview-week2/handler"
)

func main() {
	db := config.SetDb()
	defer db.Close()

	router, server := config.SetServer()
	dbHandler := handler.NewDbHandler(db)
	storeHandler := handler.NewStoreHandler(*dbHandler)

	router.GET("/branches", storeHandler.GetBranches)
	router.GET("/branches/:id", storeHandler.GetBranchById)
	router.PUT("/branches/:id", storeHandler.PutBranch)
	router.DELETE("/branches/:id", storeHandler.DeleteBranch)
	router.POST("/branches/", storeHandler.PostBranch)

	panic(server.ListenAndServe())
}