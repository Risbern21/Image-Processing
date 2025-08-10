package routes

import (
	images "image/controllers/files"

	"github.com/gofiber/fiber/v2"
)

func Images(r fiber.Router) {
	imageRouter := r.Group("/users/:id/images")

	imageRouter.Post("/", images.UploadImage)
	imageRouter.Get("/:filename", images.DownloadImage)

	imageRouter.Get("/i_id", nil)
	imageRouter.Put("/:i_id", nil)
	imageRouter.Delete("/:i_id", nil)
}
