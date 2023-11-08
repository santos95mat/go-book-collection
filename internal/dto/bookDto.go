package dto

type BookInputDTO struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Gender string `json:"gender"`
	Year   string `json:"year"`
}
