package handler

import (
	"goauthors/internal/service"
	"goauthors/internal/util"
)

type Handler struct {
	logger        util.IAppLogger
	authorService service.IAuthorService
}

func New(
	logger util.IAppLogger,
	authorService service.IAuthorService,
) *Handler {
	return &Handler{
		logger:        logger,
		authorService: authorService,
	}
}
