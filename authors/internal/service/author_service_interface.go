package service

import (
	"context"
	"goauthors/internal/dto/response"
)

type IAuthorService interface {
	FindAll(ctx context.Context) ([]*response.AuthorDto, error)
	FindByID(ctx context.Context, id int) (*response.AuthorDto, error)
}
