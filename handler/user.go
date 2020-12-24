package handler

import (
	"net/http"

	"github.com/dev-sota/echo-gorm-graphql-example/model"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	var u []*model.User

	if err := db.Find(&u).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
