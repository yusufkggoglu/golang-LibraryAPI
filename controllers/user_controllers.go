package controllers

import (
	"library/dtos"
	"library/service"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	res := service.GetAllUsers()
	return res
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Conversion err ", err)
	}
	res := service.GetUser(n)
	return res
}

func CreateUser(c *fiber.Ctx) error {
	userDto := new(dtos.CreateUserDto)
	err := c.BodyParser(userDto)
	if err != nil {
		return NewResponse(500, "Something's wrong with your input", "nil")
	}
	res := service.CreateUser(*userDto)
	return res
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal("Conversion err ", err)
	}
	res := service.DeleteUser(n)
	return res
}
