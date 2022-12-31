package middleware

import (
	"gogateway/internal/util"
)

type Middleware struct {
	logger util.IAppLogger
}

func New(logger util.IAppLogger) *Middleware {
	return &Middleware{
		logger: logger,
	}
}
