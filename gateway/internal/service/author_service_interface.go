package service

import (
	"context"
	"gogateway/internal/dto/response"
)

type IAuthorService interface {
	FindAllAuthors(ctx context.Context) (*response.Standard, error)
	FindAuthorByID(ctx context.Context, id string) (*response.Standard, error)
}
