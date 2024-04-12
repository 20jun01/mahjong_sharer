package router

import "github.com/labstack/echo/v4"

type API struct {
}

func NewAPI() API {
	return API{}
}

func SetUpRouter(e *echo.Echo, api API) {
	e.GET("/ping", nil)

	authGroup := e.Group("/auth")
	{
		authGroup.POST("/signup", nil)
	}

	// use auth
	e.Use(nil)

	e.GET("/me", nil)

	e.Logger.Fatal(e.Start(":8080"))
}
