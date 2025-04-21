package controllers

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"log/slog"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
	"github.com/nickknoddy/skyshift/pkg/processors"
)

func Transform(c *fiber.Ctx) error {

	fileName := c.Params("fileName")
	var buf bytes.Buffer

	image, err := imaging.Open(fmt.Sprintf("tmp/%s", fileName))
	if err != nil {
		slog.Error("failed to open image", "reason", err)
		return c.Status(500).SendString("Failed to process image")
	}

	imageType := processors.InferImageType(fileName)
	c.Type(imageType)

	if c.Query("w") != "" && c.Query("h") != "" {
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

		buf, err := processors.Resize(image, w, h, imaging.Lanczos, imageType)
		if err != nil {
			return c.Status(400).SendString("Failed to process image")
		}

		return c.Send(buf)
	}

	if c.Query("sharpen") != "" {
		sharpen, err := strconv.ParseFloat(c.Query("sharpen"), 64)
		if err != nil {
			slog.Error("sharpen should be an int")
			return c.Status(400).SendString("sharpen should be an int")
		}

		if sharpen > 0 {

			buf, err := processors.Sharpen(image, sharpen, imageType)
			if err != nil {
				return c.Status(400).SendString("Failed to process image")
			}

			return c.Send(buf)
		}
	}

	if c.Query("blur") != "" {

		blur, err := strconv.ParseFloat(c.Query("blur"), 64)
		if err != nil {
			slog.Error("blur should be an int")
			return c.Status(400).SendString("blur should be an int")
		}

		if blur > 0 {
			blurImage := imaging.Blur(image, blur)
			if err := jpeg.Encode(&buf, blurImage, &jpeg.Options{Quality: 50}); err != nil {
				return c.Status(500).SendString("Failed to process image")
			}

			c.Type("jpg")
		}
	}

	if c.Query("brightness") != "" {

		brightness, err := strconv.ParseFloat(c.Query("brightness"), 64)
		if err != nil {
			slog.Error("brightness should be an int")
			return c.Status(400).SendString("brightness should be an int")
		}

		if brightness > 0 {
			brightenedImage := imaging.AdjustBrightness(image, brightness)
			if err := jpeg.Encode(&buf, brightenedImage, &jpeg.Options{Quality: 50}); err != nil {
				return c.Status(500).SendString("Failed to process image")
			}

			c.Type("jpg")
		}
	}

	if c.Query("contrast") != "" {

		contrast, err := strconv.ParseFloat(c.Query("contrast"), 64)
		if err != nil {
			slog.Error("contrast should be an int")
			return c.Status(400).SendString("contrast should be an int")
		}

		if contrast > 0 {
			contrastImage := imaging.AdjustContrast(image, contrast)
			if err := jpeg.Encode(&buf, contrastImage, &jpeg.Options{Quality: 50}); err != nil {
				return c.Status(500).SendString("Failed to process image")
			}

			c.Type("jpg")
		}
	}

	if c.Query("flip") != "" {

		if c.Query("flip") == "h" {
			flippedImage := imaging.FlipH(image)
			if err := jpeg.Encode(&buf, flippedImage, &jpeg.Options{Quality: 50}); err != nil {
				return c.Status(500).SendString("Failed to process image")
			}

			c.Type("jpg")
		} else if c.Query("flip") == "v" {
			flippedImage := imaging.FlipV(image)
			if err := jpeg.Encode(&buf, flippedImage, &jpeg.Options{Quality: 50}); err != nil {
				return c.Status(500).SendString("Failed to process image")
			}

			c.Type("jpg")
		}
	}

	return c.Send(buf.Bytes())
}
