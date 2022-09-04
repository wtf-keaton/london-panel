package middleware

import (
	"github.com/gofiber/fiber/v2"
	"template/pgk/session_manager"
)

func AuthCheck(c *fiber.Ctx) error {
	if session_manager.IsAuthorized(c) == true {
		return c.Next()
	}

	if session_manager.GetUser(c).Username == "" {
		return c.Redirect("/")
	}

	return c.Render("auth", fiber.Map{})
}
