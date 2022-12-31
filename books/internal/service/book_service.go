package service

import (
	"context"
	"errors"
	"fmt"
	"gobooks/internal/dto/response"
	"gobooks/internal/repository"
	"gobooks/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type bookService struct {
	logger         util.IAppLogger
	bookRepository repository.IBookRepository
}

func NewBookService(logger util.IAppLogger, bookRepository repository.IBookRepository) IBookService {
	return &bookService{
		logger:         logger,
		bookRepository: bookRepository,
	}
}

func (bs bookService) FindAll(ctx context.Context) ([]*response.BookDto, error) {
	const scope = "service#FindAll"
	books, err := bs.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, &ErrUnknown)
	}
	result := []*response.BookDto{}
	for _, book := range books {
		result = append(result, book.ToResponseDto())
	}
	ginCtx := ctx.(*gin.Context)
	bs.logger.Info(
		"Get list of books",
		logrus.Fields{
			"id":    ginCtx.Request.Header.Get("X-Request-Id"),
			"scope": scope,
		},
	)
	return result, nil
}

func (bs bookService) FindByID(ctx context.Context, id int) (*response.BookDto, error) {
	const scope = "service#FindByID"
	book, err := bs.bookRepository.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, &repository.ErrDataNotFound) {
			return nil, fmt.Errorf("%s: %w", scope, ErrBookNotFound.SetError(err))
		}
		return nil, fmt.Errorf("%s: %w", scope, &ErrUnknown)
	}
	ginCtx := ctx.(*gin.Context)
	bs.logger.Info(
		"Get a book by its id",
		logrus.Fields{
			"id":                ginCtx.Request.Header.Get("X-Request-Id"),
			"scope":             scope,
			"requested_book_id": book.ID,
		},
	)
	return book.ToResponseDto(), nil
}
