package infrastructures

import (
	"fmt"

	"github.com/labstack/echo"
)

type AuthMiddleware interface {
	LoginMiddleware(next echo.HandlerFunc) echo.HandlerFunc
	AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type AuthMiddlewareImpl struct {
}

func NewAuthMiddleware() *AuthMiddlewareImpl {
	return &AuthMiddlewareImpl{}
}

func (a *AuthMiddlewareImpl) LoginMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Not implemented.")
		return next(c)
	}
}

func (a *AuthMiddlewareImpl) AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Not implemented.")
		return next(c)
	}
}
