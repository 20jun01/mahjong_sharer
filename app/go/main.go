package main

import (
	"fmt"
	"log"
	handler "mahjong_backend/pkg/adapter/controller"
	"mahjong_backend/pkg/infrastructures"
	"mahjong_backend/pkg/infrastructures/database"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-Authorization"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// ヘルスチェック
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	})

	authMiddleware := infrastructures.NewAuthMiddleware()
	loginUsers := e.Group("", authMiddleware.LoginMiddleware)
	_ = loginUsers.Group("/admin", authMiddleware.AdminMiddleware)

	db := database.NewClientImpl()
	defer func(db *database.ClientImpl) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	handler.NewUserController(loginUsers, db)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

func healthSkipper(c echo.Context) bool {
	return c.Path() == "/health"
}
