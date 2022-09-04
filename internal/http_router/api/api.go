package api

import (
	"github.com/gofiber/fiber/v2"
	"template/internal/models"
)

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