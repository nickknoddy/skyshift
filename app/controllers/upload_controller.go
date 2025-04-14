package controllers

import (
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		slog.Error("Error while getting file", "Reason:", err)
		return err
	}
	destination := fmt.Sprintf("./tmp/%s", file.Filename)
	if err := c.SaveFile(file, destination); err != nil {
		slog.Error("Error while saving file", "Reason:", err)
		return err
	}
	c.JSON(fiber.Map{
		"status": "success",
		"error":  nil,
		"data":   nil,
		"msg":    "File successfully uploaded",
	})

	return nil
}
