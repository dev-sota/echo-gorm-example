package handlers

import (
	"net/http"

	"github.com/dev-sota/echo-gorm-example/models"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	var u []*models.User

	if err := db.Find(&u).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
