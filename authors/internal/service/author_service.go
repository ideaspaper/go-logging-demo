package service

import (
	"context"
	"errors"
	"fmt"
	"goauthors/internal/dto/response"
	"goauthors/internal/repository"
	"goauthors/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type authorService struct {
	logger           util.IAppLogger
	authorRepository repository.IAuthorRepository
}

func NewAuthorService(logger util.IAppLogger, authorRepository repository.IAuthorRepository) IAuthorService {
	return &authorService{
		logger:           logger,
		authorRepository: authorRepository,
	}
}

func (bs authorService) FindAll(ctx context.Context) ([]*response.AuthorDto, error) {
	const scope = "service#FindAll"
	authors, err := bs.authorRepository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, &ErrUnknown)
	}
	result := []*response.AuthorDto{}
	for _, author := range authors {
		result = append(result, author.ToResponseDto())
	}
	ginCtx := ctx.(*gin.Context)
	bs.logger.Info(
		"Get list of authors",
		logrus.Fields{
			"id":    ginCtx.Request.Header.Get("X-Request-Id"),
			"scope": scope,
		},
	)
	return result, nil
}

func (bs authorService) FindByID(ctx context.Context, id int) (*response.AuthorDto, error) {
	const scope = "service#FindByID"
	author, err := bs.authorRepository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, &repository.ErrDataNotFound) {
			return nil, fmt.Errorf("%s: %w", scope, ErrAuthorNotFound.SetError(err))
		}
		return nil, fmt.Errorf("%s: %w", scope, &ErrUnknown)
	}
	ginCtx := ctx.(*gin.Context)
	bs.logger.Info(
		"Get a author by its id",
		logrus.Fields{
			"id":                  ginCtx.Request.Header.Get("X-Request-Id"),
			"scope":               scope,
			"requested_author_id": author.ID,
		},
	)
	return author.ToResponseDto(), nil
}
