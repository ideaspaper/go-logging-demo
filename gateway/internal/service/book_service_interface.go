package service

import (
	"context"
	"gogateway/internal/dto/response"
)

type IBookService interface {
	FindAllBooks(ctx context.Context) (*response.Standard, error)
	FindBookByID(ctx context.Context, id string) (*response.Standard, error)
}
