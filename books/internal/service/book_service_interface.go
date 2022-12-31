package service

import (
	"context"
	"gobooks/internal/dto/response"
)

type IBookService interface {
	FindAll(ctx context.Context) ([]*response.BookDto, error)
	FindByID(ctx context.Context, id int) (*response.BookDto, error)
}
