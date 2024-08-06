package fiber

import (
	"github.com/gofiber/fiber/v2"

	"parliamenthack-api/endpoint"
	"parliamenthack-api/type/response"
	"parliamenthack-api/util/log"
)

func Init() {
	// Initialize fiber instance
	app := fiber.New(fiber.Config{
		AppName:       "Lo Pro #parliamenthack",
		ErrorHandler:  ErrorHandler,
		Prefork:       false,
		StrictRouting: true,
	})

	// Register root endpoint
	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(response.Info("Lo Pro API ROOT"))
	})

	// Register API endpoints
	apiGroup := app.Group("api/")
	endpoint.Init(apiGroup)

	// Register not found endpoint
	app.Use(NotFoundHandler)

	// Startup
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Unable to start fiber instance", err)
	}
}
