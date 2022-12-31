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

type authorService struct {
	logger util.IAppLogger
}

func NewAuthorService(logger util.IAppLogger) IAuthorService {
	return &authorService{
		logger: logger,
	}
}

func (as authorService) FindAllAuthors(ctx context.Context) (*resp.Standard, error) {
	const scope = "service#FindAllAuthors"
	bodyDto := resp.Standard{}
	request, err := http.NewRequest(http.MethodGet, "http://authors_backend:8082/authors", nil)
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

func (as authorService) FindAuthorByID(ctx context.Context, id string) (*resp.Standard, error) {
	const scope = "service#FindAuthorByID"
	bodyDto := resp.Standard{}
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("http://authors_backend:8082/authors/%s", id),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", scope, ErrCreatingRequestFailed.SetError(err))
	}
	request.Header.Set("X-Request-Id", ctx.Value("X-Request-Id").(string))
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
