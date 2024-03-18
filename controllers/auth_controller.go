package controllers

import (
	"library/dtos"
	"library/repository"
	"library/service"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Login(c *fiber.Ctx) error {
	userDto := new(dtos.LoginUserDto)
	err := c.BodyParser(userDto)
	if err != nil {
		return NewResponse(401, "invalid input", "nil")
	}
	sessDto, res := service.Login(*userDto)
	if sessDto != nil {
		sess := GetSession(c)
		// sess.Set("user", sessDto)
		sess.Set("username", sessDto.Username)
		sess.Set("role", sessDto.Role)
		if err := sess.Save(); err != nil {
			log.Fatal(err)
		}
	}
	return res
}

func Logout(c *fiber.Ctx) error {
	sess := GetSession(c)
	if err := sess.Destroy(); err != nil {
		log.Fatal(err)
	}
	return NewResponse(200, "Logout successfuly", nil)
}

func GetSession(c *fiber.Ctx) *session.Session {
	sess, err := repository.Store.Get(c)
	if err != nil {
		panic(err)
	}
	return sess
}
