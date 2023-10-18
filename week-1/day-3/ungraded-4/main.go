package main

import (
	"ungraded-4/config"
	"ungraded-4/handler"
)

func main() {
	router, server := config.SetupServer()
	dbHandler := handler.NewDbHandler(config.ConnectDb())
	defer dbHandler.Close()

	router.PanicHandler = handler.PanicFunc
	router.GET("/heroes", dbHandler.GetHeroes)
	router.GET("/heroes/:id", dbHandler.GetHeroById)
	router.GET("/villains", dbHandler.GetVillains)
	router.GET("/villains/:id", dbHandler.GetVillainById)

	router.GET("/crimereports", dbHandler.GetCrimeReports)
	router.GET("/crimereports/:id", dbHandler.GetCrimeReportsId)
	router.POST("/crimereports", dbHandler.PostCrimeReport)
	router.PUT("/crimereports/:id", dbHandler.PutCrimeReport)
	router.DELETE("/crimereports/:id", dbHandler.DeleteCrimeReport)

	panic(server.ListenAndServe())
}