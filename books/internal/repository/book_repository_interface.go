package repository

import (
	"context"
	"gobooks/internal/model"
)

type IBookRepository interface {
	FindAll(ctx context.Context) ([]*model.Book, error)
	FindByID(ctx context.Context, id int) (*model.Book, error)
}
