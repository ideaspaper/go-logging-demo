package model

import "gobooks/internal/dto/response"

type Book struct {
	ID       int
	Title    string
	AuthorID int
}

func (b Book) ToResponseDto() *response.BookDto {
	return &response.BookDto{
		ID:       b.ID,
		Title:    b.Title,
		AuthorId: b.AuthorID,
	}
}
