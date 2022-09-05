package admin

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"template/internal/models"
	"template/pgk/memcache"
)

func CreateCheat(c *fiber.Ctx) error {
	err := models.DB.Create(&models.CheatModel{
		Name:      c.FormValue("name"),
		Status:    0,
		Creator:   c.FormValue("creator"),
		Filename:  c.FormValue("filename"),
		Process:   c.FormValue("process"),
		Anticheat: c.FormValue("anticheat"),
	}).Error

	if err != nil {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	go memcache.CheatCache.Fetch()

	return c.Redirect("/admin/cheats")
}

func UploadFile(c *fiber.Ctx) error {
	cheat := c.Params("cheat")

	id, _ := strconv.Atoi(cheat)
	cheatModel := memcache.CheatCache.ID(uint(id))

	file, err := c.FormFile("cheat")
	if err != nil {
		panic(err)
	}
	c.SaveFile(file, fmt.Sprintf("./dlls/%s", cheatModel.Filename))

	return c.Redirect("/admin/cheats")
}

func ChangeCheatStatus(c *fiber.Ctx) error {
	cheat := c.Params("cheat")

	id, _ := strconv.Atoi(cheat)
	cheatModel := memcache.CheatCache.ID(uint(id))

	if cheatModel.Status == 1 {
		models.DB.Model(&cheatModel).Update("Status", 0)
	} else if cheatModel.Status == 0 {
		models.DB.Model(&cheatModel).Update("Status", 1)
	}

	memcache.CheatCache.Fetch()

	return c.Redirect("/admin/cheats")
}

func DeleteCheat(c *fiber.Ctx) error {
	cheat := c.Params("cheat")

	id, _ := strconv.Atoi(cheat)
	cheatModel := memcache.CheatCache.ID(uint(id))

	models.DB.Delete(&cheatModel)

	memcache.CheatCache.Fetch()

	return c.Redirect("/admin/cheats")
}
