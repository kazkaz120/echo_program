package main

import (
	"github.com/kazkaz120/echo_program/handlers"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks)

	e.Start(":8080")
}
