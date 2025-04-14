package controllers

import (
	"bytes"
	"fmt"
	"image/png"
	"log/slog"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
)

func Transform(c *fiber.Ctx) error {

	fileName := c.Params("fileName")

	w, err := strconv.Atoi(c.Query("w"))
	if err != nil {
		slog.Error("width should be an int")
		return c.Status(400).SendString("width should be an int")
	}
	h, err := strconv.Atoi(c.Query("h"))
	if err != nil {
		slog.Error("height should be an int")
		return c.Status(400).SendString("height should be an int")
	}

	image, err := imaging.Open(fmt.Sprintf("tmp/%s", fileName))
	if err != nil {
		slog.Error("failed to open image")
		return c.Status(500).SendString("Failed to process image")
	}

	transformedImage := imaging.Resize(image, w, h, imaging.Lanczos)

	var buf bytes.Buffer
	if err := png.Encode(&buf, transformedImage); err != nil {
		return c.Status(500).SendString("Failed to process image")
	}

	c.Type("png")
	return c.Send(buf.Bytes())
}
