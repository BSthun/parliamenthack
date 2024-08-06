package fiber

import (
	"fmt"

	"parliamenthack-api/type/response"
)

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
		Success: false,
		Message: fmt.Sprintf("%s %s not found", c.Method(), c.Path()),
		Error:   "404_NOT_FOUND",
	})
}
