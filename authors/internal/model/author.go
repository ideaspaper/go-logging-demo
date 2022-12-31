package model

import "goauthors/internal/dto/response"

type Author struct {
	ID   int
	Name string
}

func (a Author) ToResponseDto() *response.AuthorDto {
	return &response.AuthorDto{
		ID:   a.ID,
		Name: a.Name,
	}
}
