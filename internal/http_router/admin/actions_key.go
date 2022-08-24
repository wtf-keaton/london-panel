package admin

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"template/internal/models"
	"template/pgk/generator"
	"template/pgk/memcache"
	"time"
)

func GenerateKeys(c *fiber.Ctx) error {
	keysAmount, _ := strconv.Atoi(c.FormValue("amount", "1"))
	keyHour, _ := strconv.Atoi(c.FormValue("hours", "24"))
	keyCreator := c.FormValue("creator")

	if len(keyCreator) == 0 {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	err := models.DB.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < keysAmount; i++ {
			if err := tx.Create(&models.KeyModel{
				Model:      gorm.Model{},
				Keycode:    generator.RandStringRunes(18),
				Status:     0,
				HardwareID: 0,
				Hours:      keyHour,
				EndTime:    time.Now(),
				CreatedBy:  keyCreator,
			}).Error; err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	go memcache.KeyCache.Fetch()

	return c.JSON(fiber.Map{"Status": "OK"})
}

func ClearKeyHardwareID(c *fiber.Ctx) error {
	keyCode := c.FormValue("key")
	if len(keyCode) == 0 {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	key := memcache.KeyCache.Get(keyCode)
	if key.Keycode != keyCode {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	models.DB.Model(&key).Updates(map[string]interface{}{
		"HardwareID": "",
	})

	go memcache.KeyCache.Fetch()

	return c.JSON(fiber.Map{"Status": "OK"})
}

func DeleteKey(c *fiber.Ctx) error {
	keyCode := c.FormValue("key")
	if len(keyCode) == 0 {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	key := memcache.KeyCache.Get(keyCode)
	if key.Keycode != keyCode {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	models.DB.Delete(&key)

	go memcache.KeyCache.Fetch()

	return c.JSON(fiber.Map{"Status": "OK"})
}
