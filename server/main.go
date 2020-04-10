package main

import (
	"algorithm/api"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", api.HelloWorld)

	e.Logger.Fatal(e.Start(":1323"))
}
