package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nickknoddy/skyshift/app/controllers"
)

func UploadRoute(app *fiber.App) {
	route := app.Group("/api/v1/upload")

	route.Post("", controllers.Upload)
}
