package response

type BookDto struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorId int    `json:"author_id"`
}
