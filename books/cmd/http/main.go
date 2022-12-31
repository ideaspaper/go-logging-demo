package main

import (
	"fmt"
	"gobooks/cmd/http/internal/handler"
	"gobooks/cmd/http/internal/middleware"
	"gobooks/cmd/http/internal/router"
	"gobooks/internal/model"
	"gobooks/internal/repository/memory"
	"gobooks/internal/service"
	"gobooks/internal/util"
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
		Filename:   "logs/books.log",
		MaxBackups: 10,
		MaxSize:    10,
		MaxAge:     30,
	}))
	return util.NewLogger(loggerInstance)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	logger := initLogger()
	db := []*model.Book{
		{
			ID:       1,
			Title:    "Book Title 1",
			AuthorID: 1,
		},
		{
			ID:       2,
			Title:    "Book Title 2",
			AuthorID: 2,
		},
		{
			ID:       3,
			Title:    "Book Title 3",
			AuthorID: 1,
		},
	}
	bookRepository := memory.NewBookRepository(logger, db)
	bookService := service.NewBookService(logger, bookRepository)
	handler := handler.New(logger, bookService)
	middleware := middleware.New(logger)
	router := router.New(handler, middleware)
	if err := router.Run(fmt.Sprintf(":%s", "8081")); err != nil {
		log.Fatalln(err)
	}
}
