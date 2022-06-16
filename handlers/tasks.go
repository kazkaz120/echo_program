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
