package handlers

import (
	"net/http"

	"C:\Users\syuka\simple-webapp-master\go-echo\models"
	"github.com/labstack/echo"
)

// GetTasks endpoint
func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetTasks())
}
