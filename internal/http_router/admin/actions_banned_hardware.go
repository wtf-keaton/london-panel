package admin

import (
	"github.com/gofiber/fiber/v2"
	"template/internal/models"
	"template/pgk/memcache"
)

func BanHardware(c *fiber.Ctx) error {
	models.DB.Create(&models.BannedHardware{
		HardwareID: c.FormValue("hardware"),
		Reason:     c.FormValue("reason"),
	})

	go memcache.BannedCache.Fetch()

	return c.Redirect("/admin/banned_hwids")
}

func UnbanHardware(c *fiber.Ctx) error {
	hardware := c.Params("hardware")

	hwid := memcache.BannedCache.Get(hardware)
	models.DB.Delete(&hwid, "`hardware_id` = ?", hardware)

	go memcache.BannedCache.Fetch()

	return c.Redirect("/admin/banned_hwids")
}
