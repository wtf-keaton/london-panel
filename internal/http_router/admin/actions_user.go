package admin

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"template/internal/models"
	"template/pgk/memcache"
)

func CreateUser(c *fiber.Ctx) error {
	models.DB.Create(&models.UserModel{
		Model:    gorm.Model{},
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
		Status:   c.FormValue("status"),
	})

	return c.JSON(fiber.Map{"Status": "OK"})
}

func DeleteUser(c *fiber.Ctx) error {
	user := c.FormValue("user")
	if len(user) == 0 {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	userModel := memcache.CheatCache.Get(user)
	if userModel.Name != user {
		return c.JSON(fiber.Map{"Status": "Error"})
	}

	models.DB.Delete(&userModel)

	return c.JSON(fiber.Map{"Status": "OK"})
}
