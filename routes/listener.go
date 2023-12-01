package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/pragmataW/to-do/endpoints"
	"github.com/pragmataW/to-do/middleware"
)

func Listener(app *fiber.App) {
	app.Post("/register", endpoints.HandleRegister)
	app.Post("/login", endpoints.HandleLogin)
	app.Get("/logout", middleware.RequireAuth, endpoints.HandleLogout)
	app.Post("/todos", middleware.RequireAuth, endpoints.AddTodos)
	app.Get("/todos", middleware.RequireAuth, endpoints.GetTodos)
	app.Get("/todos/:id", middleware.RequireAuth, endpoints.GetTodo)
	app.Patch("/todos/:id", middleware.RequireAuth, endpoints.UpdateTodo)
	app.Delete("/todos/:id", middleware.RequireAuth, endpoints.DeleteTodo)
}