package handler

import (
	"github.com/gin-gonic/gin"
)

func (h Handler) FindAllBooks(ctx *gin.Context) {
	result, err := h.bookService.FindAllBooks(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(
		result.Code,
		result,
	)
}

func (h Handler) FindBookByID(ctx *gin.Context) {
	paramID := ctx.Param("id")
	result, err := h.bookService.FindBookByID(ctx, paramID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(
		result.Code,
		result,
	)
}
