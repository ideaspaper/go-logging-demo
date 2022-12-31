package middleware

import (
	"errors"
	"goauthors/cmd/http/internal"
	"goauthors/internal/dto/response"
	"goauthors/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (m Middleware) ErrorHandler(ctx *gin.Context) {
	const scope = "middleware#ErrorHandler"
	ctx.Next()
	if len(ctx.Errors) == 0 {
		return
	}
	code := http.StatusInternalServerError
	body := &response.Standard{
		Code:    code,
		Message: http.StatusText(code),
		Data:    nil,
	}
	firstErr := ctx.Errors[0].Err
	reqID := ctx.Request.Header.Get("X-Request-Id")
	if errors.Is(firstErr, &service.ErrAuthorNotFound) {
		code = http.StatusNotFound
		body = &response.Standard{
			Code:    code,
			Message: http.StatusText(code),
			Data:    nil,
		}
	} else if errors.Is(firstErr, &internal.ErrCors) {
		code = http.StatusNoContent
		body = &response.Standard{
			Code:    code,
			Message: http.StatusText(http.StatusNoContent),
			Data:    nil,
		}
	} else if errors.Is(firstErr, &internal.ErrNoRoute) {
		code = http.StatusNotFound
		body = &response.Standard{
			Code:    code,
			Message: "Oops... nothing here",
			Data:    nil,
		}
	}
	logMessage := "Cannot process request"
	logFunction := m.logger.Info
	if code == http.StatusInternalServerError {
		logMessage = "Unhandled error"
		logFunction = m.logger.Error
	}
	logFunction(
		logMessage,
		logrus.Fields{
			"id":    reqID,
			"scope": scope,
			"error": firstErr.Error(),
		},
	)
	ctx.AbortWithStatusJSON(code, body)
}
