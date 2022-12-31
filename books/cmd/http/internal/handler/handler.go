package handler

import (
	"gobooks/internal/service"
	"gobooks/internal/util"
)

type Handler struct {
	logger      util.IAppLogger
	bookService service.IBookService
}

func New(
	logger util.IAppLogger,
	bookService service.IBookService,
) *Handler {
	return &Handler{
		logger:      logger,
		bookService: bookService,
	}
}
