package main

import (
	"net/http"
	"redtower/service/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database.ConnectDatabase()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
