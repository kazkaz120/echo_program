package handlers

import (
	"net/http"

	"github.com/kazkaz120/echo_program/models"
	"github.com/labstack/echo"
)

// GetTasks endpoint
func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetTasks())
}

func CreateTasks(c echo.Context) error {
	name := c.FormValue("name")
	name2 := c.FormValue("name2")
	name3 := c.FormValue("name3")
	return c.JSON(http.StatusOK, models.CreateTasks(name, name2, name3))
	//	Tasks.Create()

}
