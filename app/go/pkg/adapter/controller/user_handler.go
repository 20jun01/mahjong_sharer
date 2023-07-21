package handler

import (
	"mahjong_backend/pkg/infrastructures/database"

	"github.com/labstack/echo"
)

type UserController interface {
}

type UserControllerImpl struct {
}

func NewUserController(_ *echo.Group, db database.Client) *UserControllerImpl {
	return &UserControllerImpl{}
}
