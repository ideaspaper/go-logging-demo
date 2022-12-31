package main

import (
	"fmt"
	"goauthors/cmd/http/internal/handler"
	"goauthors/cmd/http/internal/middleware"
	"goauthors/cmd/http/internal/router"
	"goauthors/internal/model"
	"goauthors/internal/repository/memory"
	"goauthors/internal/service"
	"goauthors/internal/util"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func initLogger() util.IAppLogger {
	loggerInstance := logrus.New()
	loggerInstance.SetLevel(logrus.DebugLevel)
	loggerInstance.SetFormatter(&logrus.JSONFormatter{})
	loggerInstance.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "logs/authors.log",
		MaxBackups: 10,
		MaxSize:    10,
		MaxAge:     30,
	}))
	return util.NewLogger(loggerInstance)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	logger := initLogger()
	db := []*model.Author{
		{
			ID:   1,
			Name: "Acong",
		},
		{
			ID:   2,
			Name: "Djoko",
		},
	}
	authorRepository := memory.NewAuthorRepository(logger, db)
	authorService := service.NewAuthorService(logger, authorRepository)
	handler := handler.New(logger, authorService)
	middleware := middleware.New(logger)
	router := router.New(handler, middleware)
	if err := router.Run(fmt.Sprintf(":%s", "8082")); err != nil {
		log.Fatalln(err)
	}
}
