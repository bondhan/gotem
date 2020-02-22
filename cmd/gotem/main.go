package main

import (
	"net/http"

	"github.com/bondhan/gotem/manager"

	"github.com/bondhan/gotem/mockup"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
)

func main() {
	manager.GetContainer()

	mockup.DoPopulateData()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
