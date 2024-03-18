package controllers

import (
	"library/dtos"
	"library/service"

	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AllBringBorrowed(c *fiber.Ctx) error {
	res := service.AllBringBorrowed()
	return res
}

func BringBorrowedByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	res := service.BringBorrowedByUsername(username)
	return res
}
func BringBorrowedByBookId(c *fiber.Ctx) error {
	id := c.Params("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Conversion err ", err)
	}
	res := service.BringBorrowedByBookId(n)
	return res

}
func TooLateBringBorrowed(c *fiber.Ctx) error {
	res := service.TooLateBringBorrowed()
	return res

}

func CreateBorrow(c *fiber.Ctx) error {
	borrowDto := new(dtos.CreateBorrowDto)
	err := c.BodyParser(borrowDto)
	if err != nil {
		return NewResponse(500, "Something's wrong with your input", "nil")
	}
	res := service.CreateBorrow(*borrowDto)
	return res
}

func UpdateBorrow(c *fiber.Ctx) error {
	id := c.Params("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Conversion err ", err)
	}
	res := service.UpdateBorrow(n)
	return res
}

func DeleteBorrow(c *fiber.Ctx) error {
	id := c.Params("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Conversion err ", err)
	}
	res := service.DeleteBorrow(n)
	return res
}
