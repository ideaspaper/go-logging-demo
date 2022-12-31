package middleware

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (m Middleware) Logger(ctx *gin.Context) {
	const scope = "middleware#Logger"
	ctx.Writer = &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
	ctx.Next()
	reqID, _ := ctx.Get("X-Request-Id")
	m.logger.Info(
		"Request",
		logrus.Fields{
			"id":          reqID,
			"scope":       scope,
			"ip":          ctx.ClientIP(),
			"method":      ctx.Request.Method,
			"path":        ctx.Request.URL.Path,
			"status_code": ctx.Writer.Status(),
		},
	)
}
