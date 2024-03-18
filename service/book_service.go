package service

import (
	"context"
	"library/dtos"
	"library/errorHandler"
	"library/repository"
)

func GetAllBooks() error {
	books, err := repository.Dbs.ListBooks(context.Background())
	if err != nil {
		return &errorHandler.ErrorResponse{
			Status:  500,
			Message: "Database error",
			Data:    nil,
		}
	}
	if len(books) == 0 {
		return &errorHandler.ErrorResponse{
			Status:  404,
			Message: "Book not found",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  200,
		Message: "Book found",
		Data:    books,
	}
}

func GetBook(id int) error {
	book, err := repository.Dbs.GetBook(context.Background(), int32(id))
	if err != nil {
		return &errorHandler.ErrorResponse{
			Status:  404,
			Message: "Book not found",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  200,
		Message: "Book found",
		Data:    book,
	}
}

func CreateBook(bookDto dtos.CreateBookDto) error {
	returnedBook, err := repository.Dbs.CreateBook(context.Background(), repository.CreateBookParams(bookDto))
	if err != nil {
		return &errorHandler.ErrorResponse{
			Status:  500,
			Message: "Could not create book",
			Data:    err,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  201,
		Message: "Book has created",
		Data:    returnedBook,
	}
}

func UpdateBook(bookDto dtos.UpdateBookDto) error {
	if _, err := repository.Dbs.GetBook(context.Background(), bookDto.ID); err != nil {
		return &errorHandler.ErrorResponse{
			Status:  404,
			Message: "Book not found",
			Data:    nil,
		}
	}
	if err := repository.Dbs.UpdateBook(context.Background(), repository.UpdateBookParams(bookDto)); err != nil {
		return &errorHandler.ErrorResponse{
			Status:  422,
			Message: "Failed to update book",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  204,
		Message: "Book updated",
		Data:    nil,
	}
}

func DeleteBook(id int) error {
	if _, err := repository.Dbs.GetBook(context.Background(), int32(id)); err != nil {
		return &errorHandler.ErrorResponse{
			Status:  404,
			Message: "Book not found",
			Data:    nil,
		}
	}
	if err := repository.Dbs.DeleteBook(context.Background(), int32(id)); err != nil {
		return &errorHandler.ErrorResponse{
			Status:  422,
			Message: "Failed to delete book",
			Data:    nil,
		}
	}
	if err := repository.Dbs.DeleteBorrowByBookId(context.Background(), int32(id)); err != nil {
		return &errorHandler.ErrorResponse{
			Status:  404,
			Message: "Failed to error while deleting to borrowed for the book",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  204,
		Message: "Book deleted",
		Data:    nil,
	}
}
