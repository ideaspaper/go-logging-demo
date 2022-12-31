package handler

import (
	"gogateway/internal/util"

	"gogateway/internal/service"
)

type Handler struct {
	logger        util.IAppLogger
	authorService service.IAuthorService
	bookService   service.IBookService
}

func New(
	logger util.IAppLogger,
	authorService service.IAuthorService,
	bookService service.IBookService,
) *Handler {
	return &Handler{
		logger:        logger,
		authorService: authorService,
		bookService:   bookService,
	}
}
