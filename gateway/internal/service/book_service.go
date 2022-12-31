package service

import (
	"context"
	"encoding/json"
	"fmt"
	resp "gogateway/internal/dto/response"
	"gogateway/internal/util"
	"io"
	"net/http"
)

type bookService struct {
	logger util.IAppLogger
}

func NewBookService(logger util.IAppLogger) IBookService {
	return &bookService{
		logger: logger,
	}
}

func (bs bookService) FindAllBooks(ctx context.Context) (*resp.Standard, error) {
	const scope = "service#FindAllBooks"
	bodyDto := resp.Standard{}
	request, err := http.NewRequest(http.MethodGet, "http://books_backend:8081/books", nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, ErrCreatingRequestFailed.SetError(err))
	}
	request.Header.Add("X-Request-Id", ctx.Value("X-Request-Id").(string))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, ErrRequestFailed.SetError(err))
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, ErrReadingBodyFailed.SetError(err))
	}
	err = json.Unmarshal(body, &bodyDto)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, ErrParsingBodyFailed.SetError(err))
	}
	return &bodyDto, nil
}

func (bs bookService) FindBookByID(ctx context.Context, id string) (*resp.Standard, error) {
	const scope = "service#FindBookByID"
	bodyDto := resp.Standard{}
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://books_backend:8081/books/%s", id),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, ErrCreatingRequestFailed.SetError(err))
	}
	request.Header.Add("X-Request-Id", ctx.Value("X-Request-Id").(string))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, ErrRequestFailed.SetError(err))
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, ErrReadingBodyFailed.SetError(err))
	}
	err = json.Unmarshal(body, &bodyDto)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, ErrParsingBodyFailed.SetError(err))
	}
	return &bodyDto, nil
}
