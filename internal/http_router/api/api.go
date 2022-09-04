package api

import (
	"github.com/gofiber/fiber/v2"
	"template/internal/models"
	"template/pgk/memcache"
	"time"
)

func KeyInformation(c *fiber.Ctx) error {
	key, hwid := c.FormValue("key"), c.FormValue("hwid")

	hwidBanned := memcache.BannedCache.Get(hwid)
	if hwidBanned.HardwareID == hwid {
		return c.JSON(fiber.Map{"Status": "HardwareBanned"})
	}

	keyData := memcache.KeyCache.Get(key)
	if keyData.Status == 0 || keyData.HardwareID == "" {
		return c.JSON(fiber.Map{"Status": "Inactivated"})
	}

	if keyData.HardwareID != hwid {
		return c.JSON(fiber.Map{"Status": "Wrong HWID"})
	}

	cheatData := memcache.CheatCache.Get(keyData.Cheat)

	return c.JSON(fiber.Map{
		"Frozen":       cheatData.Status == 1,
		"AntiCheat":    cheatData.Anticheat,
		"Process":      cheatData.Process,
		"Subscription": keyData.EndTime.Format(""),
	})
}

func GetCheatFile(c *fiber.Ctx) error {
	key, hwid := c.FormValue("key"), c.FormValue("hwid")

	hwidBanned := memcache.BannedCache.Get(hwid)
	if hwidBanned.HardwareID == hwid {
		return c.JSON(fiber.Map{"Status": "HardwareBanned"})
	}

	keyData := memcache.KeyCache.Get(key)
	if keyData.Status == 0 || keyData.HardwareID == "" {
		return c.JSON(fiber.Map{"Status": "Inactivated"})
	}

	if keyData.HardwareID != hwid {
		return c.JSON(fiber.Map{"Status": "Wrong HWID"})
	}

	if keyData.EndTime.Unix() < time.Now().Unix() {
		return c.JSON(fiber.Map{"Status": "Subscription Ended"})
	}

	cheatData := memcache.CheatCache.Get(keyData.Cheat)

	return c.SendFile(cheatData.Filename)
}

func BanHardware(c *fiber.Ctx) error {
	hardware, reason := c.FormValue("hwid"), c.FormValue("reason")

	err := models.DB.Create(&models.BannedHardware{
		HardwareID: hardware,
		Reason:     reason,
	})

	if err != nil {
		c.JSON(fiber.Map{"Status": "Error"})
	}

	return c.JSON(fiber.Map{"Status": "Banned"})
}
