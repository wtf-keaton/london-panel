package admin

import (
	"github.com/gofiber/fiber/v2"
	"template/pgk/memcache"
)

func Homepage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func Users(c *fiber.Ctx) error {
	return c.Render("users", fiber.Map{
		"Users":  memcache.UserCache,
		"Cheats": memcache.CheatCache,
	})
}

func Cheats(c *fiber.Ctx) error {
	return c.Render("cheats", fiber.Map{
		"Cheats": memcache.CheatCache,
	})
}

func Keys(c *fiber.Ctx) error {
	return c.Render("keys", fiber.Map{
		"Keys":   memcache.KeyCache,
		"Cheats": memcache.CheatCache,
	})
}
