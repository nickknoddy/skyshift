package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/nickknoddy/skyshift/pkg/configs"
	"github.com/nickknoddy/skyshift/pkg/routes"
	"github.com/spf13/viper"
)

func init() {
	// Read env file
	readEnvVar(".env")
}

func readEnvVar(fileName string) {
	viper.SetConfigFile(fileName)

	err := viper.ReadInConfig()

	if err != nil {
		slog.Error("Error while reading env variables", "Reason:", err)
	}
}

func main() {

	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	routes.UploadRoute(app)
	routes.TransformRoute(app)

	app.Get("/ping", PingCheck)

	if err := app.Listen(":8000"); err != nil {
		slog.Error("Server is not starting.", "Reason:", err)
	}
}

func PingCheck(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
