package app

import (
	"newm/internal/app/postgres"
	"newm/internal/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitProject() error {
	db, err := postgres.OpenDb()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Recover())

	loggerRepo := logger.NewLoggerRepo(db)
	loggerService := logger.NewLoggerService(loggerRepo)
	loggerHandler := logger.NewLoggerHandler(loggerService)
	e = loggerHandler.InitLoggerRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))

	return nil
}
