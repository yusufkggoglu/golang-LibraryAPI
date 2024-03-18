package service

import (
	"context"
	"library/dtos"
	"library/errorHandler"
	"library/repository"
)

func GetAllUsers() error {
	users, err := repository.Dbs.ListUser(context.Background())
	if err != nil {
		return &errorHandler.ErrorResponse{
			Status:  500,
			Message: "Database error",
			Data:    nil,
		}
	}
	if len(users) == 0 {
		return &errorHandler.ErrorResponse{
			Status:  404,
			Message: "Users not found",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  200,
		Message: "Users found",
		Data:    users,
	}
}

func GetUser(id int) error {
	user, err := repository.Dbs.GetUser(context.Background(), int32(id))
	if err != nil {
		return &errorHandler.ErrorResponse{
			Status:  404,
			Message: "User not found",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  200,
		Message: "User found",
		Data:    user,
	}
}

func CreateUser(userDto dtos.CreateUserDto) error {
	returnedUser, err := repository.Dbs.CreateUser(context.Background(), repository.CreateUserParams(userDto))
	if err != nil {
		return &errorHandler.ErrorResponse{
			Status:  500,
			Message: "Could not create user",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  200,
		Message: "User has created",
		Data:    returnedUser,
	}
}

func DeleteUser(id int) error {

	if _, err := repository.Dbs.GetUser(context.Background(), int32(id)); err != nil {
		return &errorHandler.ErrorResponse{
			Status:  404,
			Message: "User not found",
			Data:    nil,
		}
	}
	if err := repository.Dbs.DeleteUser(context.Background(), int32(id)); err != nil {
		return &errorHandler.ErrorResponse{
			Status:  401,
			Message: "Failed to delete user",
			Data:    nil,
		}
	}
	if err := repository.Dbs.DeleteBorrowByUserId(context.Background(), int32(id)); err != nil {
		return &errorHandler.ErrorResponse{
			Status:  401,
			Message: "Failed to error while deleting to borrowed for the user",
			Data:    nil,
		}
	}
	return &errorHandler.ErrorResponse{
		Status:  204,
		Message: "Success",
		Data:    nil,
	}
}
