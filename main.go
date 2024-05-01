package main

import (
	"net/http"
	"redtower/service/database"
	"redtower/service/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := godotenv.Load()

	if err != nil {
		e.Logger.Fatal("Error loading .env file:", err)
	}

	database.ConnectDatabase()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.POST("/auth/login", routes.Login)
	e.POST("/auth/signup", routes.Signup)

	e.Logger.Fatal(e.Start(":8000"))
}
