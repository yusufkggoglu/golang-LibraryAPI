package controllers

import "library/errorHandler"

func NewResponse(status int, message string, data interface{}) error {
	return &errorHandler.ErrorResponse{Status: status, Message: message, Data: data}
}
