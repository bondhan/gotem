package main

import (
	"net/http"

	"github.com/bondhan/gotem/internal/initialize"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
)

func main() {
	initialize.Init()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
