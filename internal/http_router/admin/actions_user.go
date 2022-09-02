package admin

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"template/internal/models"
	"template/pgk/memcache"
	"time"
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

func LoginIn(c *fiber.Ctx) error {
	login, password := c.FormValue("Username"), c.FormValue("Password")

	userFounded := memcache.UserCache.Get(login)
	if userFounded.Username == "" || userFounded.Password != password {
		return c.Redirect("/")
	}

	c.Cookie(&fiber.Cookie{Name: "USR", Value: c.FormValue("Username"), Expires: time.Now().Add(48 * time.Hour)})
	c.Cookie(&fiber.Cookie{Name: "Authenticated", Value: "true", Expires: time.Now().Add(48 * time.Hour)})

	return c.Redirect("/admin")
}
