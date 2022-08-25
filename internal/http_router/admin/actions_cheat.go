package admin

import (
	"github.com/gofiber/fiber/v2"
	"template/internal/models"
	"template/pgk/memcache"
)

func CreateCheat(c *fiber.Ctx) error {
	err := models.DB.Create(&models.CheatModel{
		Name:     c.FormValue("name"),
		Status:   0,
		Creator:  c.FormValue("creator"),
		Filename: c.FormValue("filename"),
		Process:  c.FormValue("process"),
	}).Error

	if err != nil {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	go memcache.CheatCache.Fetch()

	return c.JSON(fiber.Map{"Status": "OK"})
}

func ChangeCheatStatus(c *fiber.Ctx) error {
	cheat := c.FormValue("cheat")
	if len(cheat) == 0 {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	cheatModel := memcache.CheatCache.Get(cheat)
	if cheatModel.Name != cheat {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	if cheatModel.Status == 0 {
		models.DB.Model(&cheatModel).Updates(map[string]interface{}{"Status": "0"})
	} else {
		models.DB.Model(&cheatModel).Updates(map[string]interface{}{"Status": "1"})
	}

	go memcache.CheatCache.Fetch()

	return c.JSON(fiber.Map{"Status": "OK"})
}

func DeleteCheat(c *fiber.Ctx) error {
	cheat := c.FormValue("cheat")
	if len(cheat) == 0 {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	cheatModel := memcache.CheatCache.Get(cheat)
	if cheatModel.Name != cheat {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	models.DB.Delete(&cheatModel)

	return c.JSON(fiber.Map{"Status": "OK"})
}
