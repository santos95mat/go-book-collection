package dto

type CreateBookDto struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Gender string `json:"gender"`
	Year   string `json:"year"`
}

type SearchBookDto struct {
	Name   string `query:"name"`
	Author string `query:"author"`
	Gender string `query:"gender"`
	Year   string `query:"year"`
}
