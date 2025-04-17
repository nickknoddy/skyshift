package controllers

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"log/slog"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
)

func Transform(c *fiber.Ctx) error {

	fileName := c.Params("fileName")
	var buf bytes.Buffer

	image, err := imaging.Open(fmt.Sprintf("tmp/%s", fileName))
	if err != nil {
		slog.Error("failed to open image", "reason", err)
		return c.Status(500).SendString("Failed to process image")
	}

	if c.Query("w") != "" && c.Query("w") != "" {
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

		transformedImage := imaging.Resize(image, w, h, imaging.Lanczos)
		if err := png.Encode(&buf, transformedImage); err != nil {
			return c.Status(500).SendString("Failed to process image")
		}

		c.Type("png")
		return c.Send(buf.Bytes())
	}

	if c.Query("sharpen") != "" {
		sharpen, err := strconv.ParseFloat(c.Query("sharpen"), 64)
		if err != nil {
			slog.Error("sharpen should be an int")
			return c.Status(400).SendString("sharpen should be an int")
		}

		if sharpen > 0 {
			img := imaging.Resize(image, image.Bounds().Dx()/2, image.Bounds().Dy()/2, imaging.Lanczos)
			slog.Info("sharpening start", "sharpen value: ", sharpen)
			sharpenedImage := imaging.Sharpen(img, sharpen)
			slog.Info("sharpening complete")
			if err := jpeg.Encode(&buf, sharpenedImage, &jpeg.Options{Quality: 50}); err != nil {
				return c.Status(500).SendString("Failed to process image")
			}

			c.Type("jpg")
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

	return c.Send(buf.Bytes())
}
