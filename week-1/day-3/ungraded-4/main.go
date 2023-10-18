package main

import (
	"ungraded-4/config"
	"ungraded-4/handler"
)

func main() {
	router, server := config.SetupServer()
	crimeHandler := handler.NewCrimeHandler(config.ConnectDb())
	defer crimeHandler.Close()

	router.PanicHandler = handler.PanicFunc
	router.GET("/crimereports", crimeHandler.GetCrimeReports)
	router.GET("/crimereports/:id", crimeHandler.GetCrimeReportsId)
	router.POST("/crimereports", crimeHandler.PostCrimeReport)
	router.PUT("/crimereports/:id", crimeHandler.PutCrimeReport)
	router.DELETE("/crimereports/:id", crimeHandler.DeleteCrimeReport)

	panic(server.ListenAndServe())
}