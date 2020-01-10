package main

import (
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo"
	"github.com/bondhan/gotem/internal/initialize"
)

func main() {
	initialize.Init()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
