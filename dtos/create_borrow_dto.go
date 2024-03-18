package dtos

type CreateBorrowDto struct {
	UserID int32  `json:"user_id"`
	BookID int32  `json:"book_id"`
	Status string `json:"status"`
}
