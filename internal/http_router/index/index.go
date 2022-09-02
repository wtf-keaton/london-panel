package index

import "github.com/gofiber/fiber/v2"

func AuthPage(c *fiber.Ctx) error {
	return c.Render("auth", fiber.Map{})
}
