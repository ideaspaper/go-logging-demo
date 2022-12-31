package middleware

import (
	"gobooks/internal/util"
)

type Middleware struct {
	logger util.IAppLogger
}

func New(logger util.IAppLogger) *Middleware {
	return &Middleware{
		logger: logger,
	}
}
