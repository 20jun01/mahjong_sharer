package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mahjong_sharer/pkg/infrastructures"
	"github.com/mahjong_sharer/pkg/router"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading env file")
	}

	e := echo.New()
	loggerWithConfig := middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: healthSkipper,
		Output:  os.Stdout,
	})
	e.Use(loggerWithConfig)
	recoverWithConfig := middleware.RecoverWithConfig(middleware.RecoverConfig{})
	e.Use(recoverWithConfig)

	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPut,
				http.MethodPost,
				http.MethodDelete,
			},
			AllowHeaders: []string{
				"*",
			},
			AllowOrigins: []string{
				"*",
			},
		}))

	_, err := infrastructures.NewGormDB()
	if err != nil {
		log.Fatal(err)
	}

	api := router.NewAPI()

	router.SetUpRouter(e, api)
}

func healthSkipper(c echo.Context) bool {
	return c.Path() == "/health"
}
