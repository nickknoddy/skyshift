package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Define a new Fiber app
	app := fiber.New()

	app.Get("/ping", PingCheck)

	if err := app.Listen(":8000"); err != nil {
		slog.Error("Server is not starting.", "Reason:", err.Error())
	}
}

func PingCheck(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
