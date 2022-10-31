package main

import (
	"github.com/BrunoKrugel/mythic-data/internal/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/char", api.CharacterData)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
