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

	models.DB.Transaction(func(tx *gorm.DB) error {
		for i := 0; i < keysAmount; i++ {
			if err := tx.Create(&models.KeyModel{
				Model:     gorm.Model{},
				Keycode:   generator.RandStringRunes(18),
				Status:    0,
				Cheat:     c.FormValue("cheat"),
				Hours:     keyHour,
				EndTime:   time.Now(),
				CreatedBy: keyCreator,
				Banned:    false,
			}).Error; err != nil {
				return err
			}
		}
		return nil
	})

	go memcache.KeyCache.Fetch()

	return c.Redirect("/admin/keys")
}

func ClearKeyHardwareID(c *fiber.Ctx) error {
	keyCode := c.Params("key")

	key := memcache.KeyCache.Get(keyCode)
	models.DB.Model(&key).Updates(map[string]interface{}{
		"HardwareID": "",
	})

	go memcache.KeyCache.Fetch()

	return c.Redirect("/admin/keys")
}

func BanKey(c *fiber.Ctx) error {
	keyCode := c.Params("key")

	key := memcache.KeyCache.Get(keyCode)
	if key.Banned {
		models.DB.Model(&key).Updates(map[string]interface{}{"Banned": 0})
	} else {
		models.DB.Model(&key).Updates(map[string]interface{}{"Banned": 1})
	}

	go memcache.KeyCache.Fetch()

	return c.Redirect("/admin/keys")
}

func DeleteKey(c *fiber.Ctx) error {
	keyCode := c.Params("key")

	key := memcache.KeyCache.Get(keyCode)
	models.DB.Delete(&key)

	go memcache.KeyCache.Fetch()

	return c.Redirect("/admin/keys")
}
