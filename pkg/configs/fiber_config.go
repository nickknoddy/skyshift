package configs

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func fiberCustomErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Send custom error page
	ctx.Status(code).SendString("Please contact the support team. Email ID: ")

	// Return from handler
	return nil
}

// FiberConfig func for configuration Fiber app.
func FiberConfig() fiber.Config {
	// Define server settings.
	appName := viper.GetString("APP_NAME")
	readTimeout := viper.GetInt("SERVER_READ_TIMEOUT")

	// Return Fiber configuration.
	return fiber.Config{
		AppName:      appName,
		ReadTimeout:  time.Second * time.Duration(readTimeout),
		ErrorHandler: fiberCustomErrorHandler,
	}
}
