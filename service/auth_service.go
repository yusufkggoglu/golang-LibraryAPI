package service

import (
	"context"
	"library/dtos"
	"library/errorHandler"
	"library/repository"
)

func Login(loginUserDto dtos.LoginUserDto) (*dtos.SessionDto, error) {
	user, _ := repository.Dbs.CheckUser(context.Background(), repository.CheckUserParams(loginUserDto))
	if (user == repository.User{}) {
		return nil, &errorHandler.ErrorResponse{
			Status:  401,
			Message: "Username or password is wrong",
			Data:    nil,
		}
	}
	return &dtos.SessionDto{Username: user.Username, Role: user.Role},
		&errorHandler.ErrorResponse{
			Status:  204,
			Message: "Succees",
			Data:    nil,
		}
}
