package api

import (
	"template/internal/models"
	"template/pgk/memcache"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ActivateKey(c *fiber.Ctx) error {
	key, hwid := c.FormValue("key"), c.FormValue("hwid")

	hwidBanned := memcache.BannedCache.Get(hwid)
	if hwidBanned.HardwareID == hwid {
		keyBanned := memcache.KeyCache.Get(key)
		models.DB.Model(&keyBanned).Updates(map[string]interface{}{"Banned": 1})

		go memcache.KeyCache.Fetch()
		return c.JSON(fiber.Map{"Status": "Hardware Banned"})
	}

	keyData := memcache.KeyCache.Get(key)
	if keyData.Status != 0 || keyData.HardwareID != "" {
		return c.JSON(fiber.Map{"Status": "Already Activated"})
	}

	models.DB.Model(&keyData).Updates(map[string]interface{}{
		"HardwareID": hwid,
		"EndTime":    keyData.EndTime.Add(time.Duration(keyData.Hours) * time.Hour * 24),
		"Status":     1,
	})

	go memcache.KeyCache.Fetch()

	return c.JSON(fiber.Map{"Status": "Activated"})
}

func CheckKey(c *fiber.Ctx) error {
	key, hwid := c.FormValue("key"), c.FormValue("hwid")

	hwidBanned := memcache.BannedCache.Get(hwid)
	if hwidBanned.HardwareID == hwid {
		keyBanned := memcache.KeyCache.Get(key)
		models.DB.Model(&keyBanned).Updates(map[string]interface{}{"Banned": 1})

		go memcache.KeyCache.Fetch()
		return c.JSON(fiber.Map{"Status": "Hardware Banned"})
	}

	keyData := memcache.KeyCache.Get(key)
	if keyData.HardwareID != hwid {
		return c.JSON(fiber.Map{"Status": "Wrong HWID"})
	}

	if keyData.HardwareID == "" {
		models.DB.Model(&keyData).Updates(map[string]interface{}{"HardwareID": hwid})
		go memcache.KeyCache.Fetch()

	}

	if keyData.EndTime.Unix() < time.Now().Unix() {
		models.DB.Model(&keyData).Updates(map[string]interface{}{"Status": 2})

		go memcache.KeyCache.Fetch()

		return c.JSON(fiber.Map{"Status": "Expired"})
	}

	return c.JSON(fiber.Map{"Status": "OK"})
}

func KeyInformation(c *fiber.Ctx) error {
	key, hwid := c.FormValue("key"), c.FormValue("hwid")

	hwidBanned := memcache.BannedCache.Get(hwid)
	if hwidBanned.HardwareID == hwid {
		keyBanned := memcache.KeyCache.Get(key)
		models.DB.Model(&keyBanned).Updates(map[string]interface{}{"Banned": 1})

		go memcache.KeyCache.Fetch()

		return c.JSON(fiber.Map{"Status": "Hardware Banned"})
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
		keyBanned := memcache.KeyCache.Get(key)
		models.DB.Model(&keyBanned).Updates(map[string]interface{}{"Banned": 1})
		go memcache.KeyCache.Fetch()

		return c.JSON(fiber.Map{"Status": "Hardware Banned"})
	}

	keyData := memcache.KeyCache.Get(key)
	if keyData.Status == 0 || keyData.HardwareID == "" {
		return c.JSON(fiber.Map{"Status": "Inactivated"})
	}

	if keyData.HardwareID != hwid {
		return c.JSON(fiber.Map{"Status": "Wrong HWID"})
	}

	if keyData.EndTime.Unix() < time.Now().Unix() {
		models.DB.Model(&keyData).Updates(map[string]interface{}{"Status": 2})

		go memcache.KeyCache.Fetch()

		return c.JSON(fiber.Map{"Status": "Expired"})
	}

	cheatData := memcache.CheatCache.Get(keyData.Cheat)

	return c.SendFile(cheatData.Filename)
}

func GetDriverFile(c *fiber.Ctx) error {
	key, hwid := c.FormValue("key"), c.FormValue("hwid")

	hwidBanned := memcache.BannedCache.Get(hwid)
	if hwidBanned.HardwareID == hwid {
		keyBanned := memcache.KeyCache.Get(key)
		models.DB.Model(&keyBanned).Updates(map[string]interface{}{"Banned": 1})
		go memcache.KeyCache.Fetch()

		return c.JSON(fiber.Map{"Status": "Hardware Banned"})
	}

	keyData := memcache.KeyCache.Get(key)
	if keyData.Status == 0 || keyData.HardwareID == "" {
		return c.JSON(fiber.Map{"Status": "Inactivated"})
	}

	if keyData.HardwareID != hwid {
		return c.JSON(fiber.Map{"Status": "Wrong HWID"})
	}

	if keyData.EndTime.Unix() < time.Now().Unix() {
		models.DB.Model(&keyData).Updates(map[string]interface{}{"Status": 2})

		go memcache.KeyCache.Fetch()

		return c.JSON(fiber.Map{"Status": "Expired"})
	}

	return c.SendFile("driver/driver.sys")
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
