package endpoint

import (
	"github.com/gofiber/fiber/v2"

	uploadEndpoint "parliamenthack-api/endpoint/upload"
)

func Init(router fiber.Router) {
	upload := router.Group("/upload")
	upload.Post("/pdfextract", uploadEndpoint.UpdateHandler)

}
