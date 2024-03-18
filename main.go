package main

import (
	"encoding/gob"
	"library/dtos"
	"library/errorHandler"
	"library/repository"
	"library/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/lib/pq"
)

func main() {
	defer repository.Conn.Close()
	repository.Connect()
	app := fiber.New(fiber.Config{ErrorHandler: errorHandler.ErrorHandler})

	gob.Register(dtos.SessionDto{})
	store := session.New()
	routers.Route(app, store)
	app.Listen(":3000")
}
