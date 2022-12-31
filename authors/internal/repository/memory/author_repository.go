package memory

import (
	"context"
	"fmt"
	"goauthors/internal/model"
	"goauthors/internal/repository"
	"goauthors/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type authorRepository struct {
	logger util.IAppLogger
	db     []*model.Author
}

func NewAuthorRepository(logger util.IAppLogger, db []*model.Author) repository.IAuthorRepository {
	return &authorRepository{
		logger: logger,
		db:     db,
	}
}

func (b authorRepository) FindAll(ctx context.Context) ([]*model.Author, error) {
	const scope = "repository#FindAll"
	ginCtx := ctx.(*gin.Context)
	b.logger.Debug(
		"Get list of authors from repository",
		logrus.Fields{
			"id":    ginCtx.Request.Header.Get("X-Request-Id"),
			"scope": scope,
		},
	)
	return b.db, nil
}

func (b authorRepository) FindByID(ctx context.Context, id int) (*model.Author, error) {
	const scope = "repository#FindByID"
	for _, author := range b.db {
		if author.ID == id {
			ginCtx := ctx.(*gin.Context)
			b.logger.Debug(
				"Get an author by its id from repository",
				logrus.Fields{
					"id":                  ginCtx.Request.Header.Get("X-Request-Id"),
					"scope":               scope,
					"requested_author_id": author.ID,
				},
			)
			return author, nil
		}
	}
	return nil, fmt.Errorf("%s: %w", scope, &repository.ErrDataNotFound)
}
