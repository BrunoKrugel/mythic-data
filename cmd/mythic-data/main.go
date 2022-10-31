package main

import (
	"log"

	"github.com/BrunoKrugel/mythic-data/internal/api"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/char", api.CharacterData)
	e.GET("/leader", api.LeaderboardHorde)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
