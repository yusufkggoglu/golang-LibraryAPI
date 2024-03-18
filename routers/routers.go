package routers

import (
	"library/controllers"
	"library/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Route(app *fiber.App, store *session.Store) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user", middleware.AuthMiddleware)
	user.Get("/", middleware.AdminAuth, controllers.GetAllUsers)
	user.Get("/:id", middleware.AdminAuth, controllers.GetUser)
	user.Post("/", middleware.AdminAuth, controllers.CreateUser)
	user.Delete("/:id", middleware.AdminAuth, controllers.DeleteUser)

	book := v1.Group("/book", middleware.AuthMiddleware)
	book.Get("/", controllers.GetAllBooks)
	book.Get("/:id", controllers.GetBook)
	book.Post("/", controllers.CreateBook)
	book.Put("/:id", middleware.AdminAuth, controllers.UpdateBook)
	book.Delete("/:id", middleware.AdminAuth, controllers.DeleteBook)

	borrow := v1.Group("/borrow", middleware.AuthMiddleware)
	borrow.Get("/", middleware.AdminAuth, controllers.AllBringBorrowed)
	borrow.Get("/BringBorrowedByUserName/:username", middleware.AdminAuth, controllers.BringBorrowedByUsername)
	borrow.Get("/BringBorrowedByBookId/:id", middleware.AdminAuth, controllers.BringBorrowedByBookId)
	borrow.Get("/TooLateBringBorrowed", middleware.AdminAuth, controllers.TooLateBringBorrowed)
	borrow.Post("/", controllers.CreateBorrow)
	borrow.Put("/:id", middleware.AdminAuth, controllers.UpdateBorrow)
	borrow.Delete("/:id", middleware.AdminAuth, controllers.DeleteBorrow)

	auth := v1.Group("/auth")
	auth.Get("/login", controllers.Login)
	auth.Get("/logout", controllers.Logout)
}
