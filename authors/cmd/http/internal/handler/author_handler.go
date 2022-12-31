package handler

import (
	"goauthors/internal/dto/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) FindAll(ctx *gin.Context) {
	authors, err := h.authorService.FindAll(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		&response.Standard{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    authors,
		},
	)
}

func (h Handler) FindByID(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.Error(err)
		return
	}
	author, err := h.authorService.FindByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		&response.Standard{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    author,
		},
	)
}
