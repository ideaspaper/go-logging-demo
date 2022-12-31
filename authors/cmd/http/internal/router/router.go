package router

import (
	"goauthors/cmd/http/internal/handler"
	"goauthors/cmd/http/internal/middleware"

	"github.com/gin-gonic/gin"
)

func New(h *handler.Handler, m *middleware.Middleware) *gin.Engine {
	r := gin.New()
	r.Use(m.Logger, m.ErrorHandler, m.CORSMiddleware)
	r.GET("/authors", h.FindAll)
	r.GET("/authors/:id", h.FindByID)
	r.NoRoute(h.NoRoute)
	return r
}
