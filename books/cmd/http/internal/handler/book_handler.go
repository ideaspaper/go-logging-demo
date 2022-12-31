package handler

import (
	"gobooks/internal/dto/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) FindAll(ctx *gin.Context) {
	books, err := h.bookService.FindAll(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		&response.Standard{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    books,
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
	book, err := h.bookService.FindByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(
		http.StatusOK,
		&response.Standard{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    book,
		},
	)
}
