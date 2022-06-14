package main

import (
	"C:\Users\syuka\simple-webapp-master\go-echo\handlers"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks)

	e.Start(":8080")
}
