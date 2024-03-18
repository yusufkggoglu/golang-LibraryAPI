package service

import (
	"context"
	"library/dtos"
	"library/errorHandler"
	"library/repository"
)

func AllBringBorrowed() error {

	borrows, err := repository.Dbs.ListBorrow(context.Background())
	if err != nil {
		return &errorHandler.ErrorResponse{
			Status:  500,
			Message: "Database error",
			Data:    nil,
		}
	}
	if len(borrows) == 0 {
		return &errorHandler.ErrorResponse{
			Status:  404,
			Message: "Borrow not found",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  200,
		Message: "Borrow found",
		Data:    borrows,
	}
}

func BringBorrowedByUsername(username string) error {
	borrow, err := repository.Dbs.BringBorrowedByUserName(context.Background(), string(username))
	if err != nil {
		return &errorHandler.ErrorResponse{
			Status:  500,
			Message: "Database error",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  200,
		Message: "Borrow found",
		Data:    borrow,
	}
}
func BringBorrowedByBookId(id int) error {
	borrow, err := repository.Dbs.BringBorrowedByBookId(context.Background(), int32(id))
	if err != nil {
		return &errorHandler.ErrorResponse{
			Status:  500,
			Message: "Database error",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  200,
		Message: "Borrows found",
		Data:    borrow,
	}
}

func TooLateBringBorrowed() error {
	borrow, err := repository.Dbs.TooLateBringBorrowed(context.Background(), "true")
	if err != nil {
		return &errorHandler.ErrorResponse{
			Status:  500,
			Message: "Database error",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  200,
		Message: "Too Late Active Borrow Found",
		Data:    borrow,
	}
}

func CreateBorrow(borrowDto dtos.CreateBorrowDto) error {
	check, _ := repository.Dbs.BringActiveBorrowByUser(context.Background(), repository.BringActiveBorrowByUserParams{
		UserID: borrowDto.UserID,
		Status: "true",
	})
	if !(check == repository.UserBook{}) {
		return &errorHandler.ErrorResponse{
			Status:  401,
			Message: "Active borrow found",
			Data:    nil,
		}
	}
	returnedBorrow, err := repository.Dbs.CreateBorrow(context.Background(), repository.CreateBorrowParams(borrowDto))
	if err != nil {
		return &errorHandler.ErrorResponse{
			Status:  500,
			Message: "Could not create borrow",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  201,
		Message: "Borrow has created",
		Data:    returnedBorrow,
	}
}

func UpdateBorrow(id int) error {
	if _, err := repository.Dbs.BringBorrowedById(context.Background(), int32(id)); err != nil {
		return &errorHandler.ErrorResponse{
			Status:  404,
			Message: "Borrow not found",
			Data:    nil,
		}
	}
	if err := repository.Dbs.UpdateStatusBorrow(context.Background(), int32(id)); err != nil {
		return &errorHandler.ErrorResponse{
			Status:  401,
			Message: "Failed to update borrow",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  200,
		Message: "Borrow updated",
		Data:    nil,
	}
}

func DeleteBorrow(id int) error {

	if _, err := repository.Dbs.BringBorrowedById(context.Background(), int32(id)); err != nil {
		return &errorHandler.ErrorResponse{
			Status:  404,
			Message: "Borrow not found",
			Data:    nil,
		}
	}
	if err := repository.Dbs.DeleteBorrow(context.Background(), int32(id)); err != nil {
		return &errorHandler.ErrorResponse{
			Status:  401,
			Message: "Failed to delete borrow",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  204,
		Message: "Success",
		Data:    nil,
	}
}
