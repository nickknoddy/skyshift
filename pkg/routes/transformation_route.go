package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nickknoddy/skyshift/app/controllers"
)

func TransformRoute(app *fiber.App) {
	route := app.Group("/api/v1/transform")

	route.Get(":fileName", controllers.Transform)
}
