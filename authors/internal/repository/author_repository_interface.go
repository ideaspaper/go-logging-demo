package repository

import (
	"context"
	"goauthors/internal/model"
)

type IAuthorRepository interface {
	FindAll(ctx context.Context) ([]*model.Author, error)
	FindByID(ctx context.Context, id int) (*model.Author, error)
}
