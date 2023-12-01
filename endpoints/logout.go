package endpoints

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HandleLogout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name: "Authentication",
		MaxAge: -1,
	})
	return c.SendStatus(http.StatusOK)
}