package app

import (
	"fmt"
	"newm/internal/app/db/postgres"
	"newm/internal/app/db/redis"
	"newm/internal/logger"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitProject() error {
	dbPsql, err := postgres.OpenDb()
	if err != nil {
		panic(err)
	}

	// dbMongo, err := mongo.OpenDb()
	// if err != nil{
	// 	panic(err)
	// }

	dbRedis, err := redis.OpenDb()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Recover())

	loggerRepoPsql := logger.NewLoggerRepo(dbPsql)
	//loggerRepoMongo := logger.NewLoggerRepoMongo(dbMongo)
	loggerRepoRedis := logger.NewLoggerRepoRedis(dbRedis)
	loggerService := logger.NewLoggerService(loggerRepoPsql, nil, loggerRepoRedis)
	loggerHandler := logger.NewLoggerHandler(loggerService)
	e = loggerHandler.InitLoggerRoutes(e)

	for _, value := range e.Routes(){
		color.New(color.FgRed, color.Bold).Println(fmt.Sprintf("\n\t%s : %s", value.Path, value.Method))
	}

	e.Logger.Fatal(e.Start(":8080"))

	return nil
}
