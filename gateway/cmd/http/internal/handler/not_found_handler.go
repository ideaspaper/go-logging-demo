package handler

import (
	"gogateway/cmd/http/internal"

	"github.com/gin-gonic/gin"
)

func (h Handler) NoRoute(ctx *gin.Context) {
	ctx.Error(&internal.ErrNoRoute)
}
