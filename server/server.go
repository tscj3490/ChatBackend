package server

import (
	"fmt"
	"net/http"

	"../api"
	"../config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Run echo framework...
func Run() {
	printLogo()

	ch := make(chan struct{})
	startServer(ch)
	<-ch
}

func startServer(exitCh chan struct{}) {
	e := echo.New()

	// Global middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Start Route API
	api.RouteAPI(e)

	// Start File server
	go startFileServer()

	if err := e.Start(":" + config.Port); err != nil {
		e.Logger.Fatal(err)
	}

	exitCh <- struct{}{}
}

func startFileServer() {
	println("startFileServer")
	handler := http.FileServer(http.Dir("./upload"))
	if err := http.ListenAndServe(":"+config.FilePort, handler); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func printLogo() {
	println("*********************************************")
	println("*-------------------------------------------*")
	println(`*             _______      ___               `)
	println(`*            /  _____|    /   \              `)
	println(`*           |  |  __     /  ^  \             `)
	println(`*           |  | |_ |   /  /_\  \            `)
	println(`*           |  |__| |  /  _____  \           `)
	println(`*            \______| /__/     \__\          `)
	println("*-------------------------------------------*")
	println("* Author: John Won")
	println("* Version: 1.0.0")
	println("* Host:", config.HostURL)
	println("* Port:", config.Port)
	println("*********************************************")
}
