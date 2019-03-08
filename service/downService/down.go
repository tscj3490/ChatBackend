package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.File("D04.png")
	})

	fmt.Println(e.Start(":3032"))
}
