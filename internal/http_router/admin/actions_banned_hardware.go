package admin

import (
	"github.com/gofiber/fiber/v2"
	"template/internal/models"
	"template/pgk/memcache"
)

func BanHardware(c *fiber.Ctx) error {
	err := models.DB.Create(&models.BannedHardware{
		HardwareID: c.FormValue("hardware"),
		Reason:     c.FormValue("reason"),
	}).Error

	if err != nil {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	go memcache.BannedCache.Fetch()

	return c.JSON(fiber.Map{"Status": "OK"})
}

func UnbanHardware(c *fiber.Ctx) error {
	hardware := c.FormValue("hardware")
	if len(hardware) == 0 {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	hwid := memcache.KeyCache.Get(hardware)
	if hwid.Keycode != hardware {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	models.DB.Delete(&hwid)

	go memcache.BannedCache.Fetch()

	return c.JSON(fiber.Map{"Status": "OK"})
}
