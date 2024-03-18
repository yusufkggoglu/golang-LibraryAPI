package middleware

import (
	"library/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	sess := controllers.GetSession(c)
	username := sess.Get("username")
	if username == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}

func AdminAuth(c *fiber.Ctx) error {
	sess := controllers.GetSession(c)
	role := sess.Get("role")
	if role != "Admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
