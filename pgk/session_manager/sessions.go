package session_manager

import (
	"github.com/gofiber/fiber/v2"
	"template/internal/models"
)

func GetUser(c *fiber.Ctx) models.UserModel {
	username := c.Cookies("USR")
	user := models.UserModel{}

	if IsAuthorized(c) == true {
		models.DB.Find(&user, "Username = ?", username)
		if user.Username == "" {
			return models.UserModel{}
		} else {
			return user
		}
	}

	return models.UserModel{}
}

func IsAuthorized(c *fiber.Ctx) bool {
	Authorized := c.Cookies("Authenticated", "false")
	if Authorized == "true" {
		return true
	}
	return false
}
