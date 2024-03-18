package controllers

import (
	"library/dtos"
	"library/service"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(c *fiber.Ctx) error {
	res := service.GetAllBooks()
	return res
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Conversion err ", err)
	}
	res := service.GetBook(n)
	return res
}

func CreateBook(c *fiber.Ctx) error {
	bookDto := new(dtos.CreateBookDto)
	err := c.BodyParser(bookDto)
	if err != nil {
		return NewResponse(500, "Something's wrong with your input", nil)
	}
	res := service.CreateBook(*bookDto)
	return res
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Conversion err ", err)
	}
	bookDto := new(dtos.UpdateBookDto)
	bookDto.ID = int32(n)
	err = c.BodyParser(bookDto)
	if err != nil {
		return NewResponse(500, "Something's wrong with your input", nil)
	}
	res := service.UpdateBook(*bookDto)
	return res
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Conversion err ", err)
	}
	res := service.DeleteBook(n)
	return res
}
