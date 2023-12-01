package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pragmataW/to-do/middleware"
)

func Listener(app *fiber.App) {
	app.Post("/register", HandleRegister)
	app.Post("/login", HandleLogin)
	app.Get("/logout", middleware.RequireAuth, HandleLogout)
	app.Post("/todos", middleware.RequireAuth, AddTodos)
	app.Get("/todos", middleware.RequireAuth, GetTodos)
	app.Get("/todos/:id", middleware.RequireAuth, GetTodo)
	app.Patch("/todos/:id", middleware.RequireAuth, UpdateTodo)
	app.Delete("/todos/:id", middleware.RequireAuth, DeleteTodo)
}