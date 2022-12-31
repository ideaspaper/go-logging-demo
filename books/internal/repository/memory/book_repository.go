package memory

import (
	"context"
	"fmt"
	"gobooks/internal/model"
	"gobooks/internal/repository"
	"gobooks/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bookRepository struct {
	logger util.IAppLogger
	db     []*model.Book
}

func NewBookRepository(logger util.IAppLogger, db []*model.Book) repository.IBookRepository {
	return &bookRepository{
		logger: logger,
		db:     db,
	}
}

func (b bookRepository) FindAll(ctx context.Context) ([]*model.Book, error) {
	const scope = "repository#FindAll"
	ginCtx := ctx.(*gin.Context)
	b.logger.Debug(
		"Get list of books from repository",
		logrus.Fields{
			"id":    ginCtx.Request.Header.Get("X-Request-Id"),
			"scope": scope,
		},
	)
	return b.db, nil
}

func (b bookRepository) FindByID(ctx context.Context, id int) (*model.Book, error) {
	const scope = "repository#FindByID"
	for _, book := range b.db {
		if book.ID == id {
			ginCtx := ctx.(*gin.Context)
			b.logger.Debug(
				"Get a book by its id from repository",
				logrus.Fields{
					"id":                ginCtx.Request.Header.Get("X-Request-Id"),
					"scope":             scope,
					"requested_book_id": book.ID,
				},
			)
			return book, nil
		}
	}
	return nil, fmt.Errorf("%s: %w", scope, &repository.ErrDataNotFound)
}
