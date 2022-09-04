package admin

import (
	"github.com/gofiber/fiber/v2"
	"template/internal/models"
	"template/pgk/memcache"
	"template/pgk/session_manager"
)

func Homepage(c *fiber.Ctx) error {
	user := session_manager.GetUser(c)
	var keys []models.KeyModel
	var keysBanned []models.KeyModel
	var Actived int

	if user.Status == "seller" {
		keys = memcache.KeyCache.SellerNotBanned(user.Username)
		keysBanned = memcache.KeyCache.SellerBanned(user.Username)
		Actived = memcache.KeyCache.SellerActived(user.Username)
	} else {
		keys = memcache.KeyCache.NotBanned()
		keysBanned = memcache.KeyCache.Banned()
		Actived = memcache.KeyCache.Actived()
	}

	return c.Render("index", fiber.Map{
		"Created": len(keys),
		"Banned":  len(keysBanned),
		"Actived": Actived,
	})
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
	user := session_manager.GetUser(c)
	var keys []models.KeyModel

	if user.Status == "seller" {
		keys = memcache.KeyCache.SellerNotBanned(user.Username)
	} else {
		keys = memcache.KeyCache.NotBanned()
	}

	return c.Render("keys", fiber.Map{
		"Keys":   keys,
		"Cheats": memcache.CheatCache,
	})
}

func KeysBanned(c *fiber.Ctx) error {
	user := session_manager.GetUser(c)
	var keys []models.KeyModel

	if user.Status == "seller" {
		keys = memcache.KeyCache.SellerBanned(user.Username)
	} else {
		keys = memcache.KeyCache.Banned()
	}

	return c.Render("keys", fiber.Map{
		"Keys":   keys,
		"Cheats": memcache.CheatCache,
	})
}
