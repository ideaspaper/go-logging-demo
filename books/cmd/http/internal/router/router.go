package router

import (
	"gobooks/cmd/http/internal/handler"
	"gobooks/cmd/http/internal/middleware"

	"github.com/gin-gonic/gin"
)

func New(h *handler.Handler, m *middleware.Middleware) *gin.Engine {
	r := gin.New()
	r.Use(m.Logger, m.ErrorHandler, m.CORSMiddleware)
	r.GET("/books", h.FindAll)
	r.GET("/books/:id", h.FindByID)
	r.NoRoute(h.NoRoute)
	return r
}
