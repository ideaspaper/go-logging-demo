package router

import (
	"gogateway/cmd/http/internal/handler"
	"gogateway/cmd/http/internal/middleware"

	"github.com/gin-gonic/gin"
)

func New(h *handler.Handler, m *middleware.Middleware) *gin.Engine {
	r := gin.New()
	r.Use(m.Logger, m.ErrorHandler, m.CORSMiddleware, m.RequestID)
	r.GET("/books", h.FindAllBooks)
	r.GET("/books/:id", h.FindBookByID)
	r.GET("/authors", h.FindAllAuthors)
	r.GET("/authors/:id", h.FindAuthorByID)
	r.NoRoute(h.NoRoute)
	return r
}
