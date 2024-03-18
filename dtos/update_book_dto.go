package dtos

type UpdateBookDto struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}
