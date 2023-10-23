package main

import (
	"preview-week2-gin/config"
	"preview-week2-gin/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetDb()
	defer db.Close()

	dbHandler := handler.NewDbHandler(db)
	storeHandler := handler.NewStoreHandler(*dbHandler)

	router := gin.Default()
	branch := router.Group("/branches")
	{
		branch.GET("", storeHandler.GetBranches)
		branch.POST("", storeHandler.PostBranch)
		branch.GET("/:id", storeHandler.GetBranchById)
		branch.PUT("/:id", storeHandler.PutBranch)
		branch.DELETE("/:id", storeHandler.DeleteBranch)
	}

	panic(router.Run())
}