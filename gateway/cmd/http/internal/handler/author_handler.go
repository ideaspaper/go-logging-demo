package handler

import (
	"github.com/gin-gonic/gin"
)

func (h Handler) FindAllAuthors(ctx *gin.Context) {
	result, err := h.authorService.FindAllAuthors(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(
		result.Code,
		result,
	)
}

func (h Handler) FindAuthorByID(ctx *gin.Context) {
	paramID := ctx.Param("id")
	result, err := h.authorService.FindAuthorByID(ctx, paramID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(
		result.Code,
		result,
	)
}
