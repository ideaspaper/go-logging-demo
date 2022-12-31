package main

import (
	"fmt"
	"gogateway/cmd/http/internal/handler"
	"gogateway/cmd/http/internal/middleware"
	"gogateway/cmd/http/internal/router"
	"gogateway/internal/service"
	"gogateway/internal/util"
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
		Filename:   "logs/gateway.log",
		MaxBackups: 10,
		MaxSize:    10,
		MaxAge:     30,
	}))
	return util.NewLogger(loggerInstance)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	logger := initLogger()
	authorService := service.NewAuthorService(logger)
	bookService := service.NewBookService(logger)
	handler := handler.New(logger, authorService, bookService)
	middleware := middleware.New(logger)
	router := router.New(handler, middleware)
	if err := router.Run(fmt.Sprintf(":%s", "8080")); err != nil {
		log.Fatalln(err)
	}
}
