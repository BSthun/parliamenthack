package upload

import "github.com/gofiber/fiber/v2"

func UpdateHandler(c *fiber.Ctx) error {
	return c.SendString("UpdateHandler")
}
