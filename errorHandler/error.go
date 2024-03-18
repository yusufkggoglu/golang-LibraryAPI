package errorHandler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *ErrorResponse) Error() string {
	return r.Message
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	base := &ErrorResponse{}
	if errors.As(err, &base) {
		return c.Status(err.(*ErrorResponse).Status).JSON(err)
	}
	fiberErr := &fiber.Error{}
	if errors.As(err, &fiberErr) {
		if fiberErr.Code == 404 {
			return c.Status(404).JSON(&ErrorResponse{
				Status:  404,
				Message: "Data not found",
			})
		} else {
			return c.Status(err.(*fiber.Error).Code).JSON(&ErrorResponse{
				Status:  err.(*fiber.Error).Code,
				Message: err.(*fiber.Error).Message,
			})
		}
	}

	return c.Status(500).JSON(&ErrorResponse{
		Status:  500,
		Message: "Internal Server Error (Unknown)", //ileride belki buna da key i18n keyi yazılabilir.
		Data:    err.Error(),
	})
}
