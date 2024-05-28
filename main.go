package main

import (
	"golang-database/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Middleware

	// Routes
	e.POST("/criminal_records", database.CreateCriminalRecord)
	e.GET("/criminal_records", database.GetCriminalRecords)
	e.GET("/criminal_records/:id", database.GetCriminalRecord)
	e.PUT("/criminal_records/:id", database.UpdateCriminalRecord)
	e.DELETE("/criminal_records/:id", database.DeleteCriminalRecord)

	e.POST("/prisons", database.CreatePrison)
	e.GET("/prisons", database.GetPrisons)
	e.GET("/prisons/:id", database.GetPrison)
	e.PUT("/prisons/:id", database.UpdatePrison)
	e.DELETE("/prisons/:id", database.DeletePrison)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
