package controllers

import (
	"bytes"
	"fmt"
	"image/color"
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
			buf, err := processors.Blur(image, blur, imageType)
			if err != nil {
				return c.Status(400).SendString("Failed to process image")
			}

			return c.Send(buf)
		}
	}

	if c.Query("brightness") != "" {

		brightness, err := strconv.ParseFloat(c.Query("brightness"), 64)
		if err != nil {
			slog.Error("brightness should be an int")
			return c.Status(400).SendString("brightness should be an int")
		}

		if brightness > 0 {
			buf, err := processors.Brightness(image, brightness, imageType)
			if err != nil {
				return c.Status(400).SendString("Failed to process image")
			}

			return c.Send(buf)
		}
	}

	if c.Query("contrast") != "" {

		contrast, err := strconv.ParseFloat(c.Query("contrast"), 64)
		if err != nil {
			slog.Error("contrast should be an int")
			return c.Status(400).SendString("contrast should be an int")
		}

		if contrast > 0 {
			buf, err := processors.Contrast(image, contrast, imageType)
			if err != nil {
				return c.Status(400).SendString("Failed to process image")
			}

			return c.Send(buf)
		}
	}

	if c.Query("flip") != "" {

		if c.Query("flip") == "h" {
			buf, err := processors.FlipHorizontal(image, imageType)
			if err != nil {
				return c.Status(400).SendString("Failed to process image")
			}

			return c.Send(buf)
		} else if c.Query("flip") == "v" {
			buf, err := processors.FlipVertical(image, imageType)
			if err != nil {
				return c.Status(400).SendString("Failed to process image")
			}

			return c.Send(buf)
		}
	}

	if c.Query("grayscale") != "" {

		buf, err := processors.GrayScale(image, imageType)
		if err != nil {
			return c.Status(400).SendString("Failed to process image")
		}

		return c.Send(buf)
	}

	if c.Query("crop-center-w") != "" && c.Query("crop-center-h") != "" {

		w, err := strconv.Atoi(c.Query("crop-center-w"))
		if err != nil {
			slog.Error("width should be an int")
			return c.Status(400).SendString("width should be an int")
		}
		h, err := strconv.Atoi(c.Query("crop-center-h"))
		if err != nil {
			slog.Error("height should be an int")
			return c.Status(400).SendString("height should be an int")
		}

		buf, err := processors.CropCenter(image, w, h, imageType)
		if err != nil {
			return c.Status(400).SendString("Failed to process image")
		}

		return c.Send(buf)
	}

	if c.Query("rotate-a") != "" {
		a, err := strconv.ParseFloat(c.Query("rotate-a"), 64)
		if err != nil {
			slog.Error("angle should be an int")
			return c.Status(400).SendString("angle should be an int")
		}

		buf, err := processors.Rotate(image, a, color.Transparent, imageType)
		if err != nil {
			return c.Status(400).SendString("Failed to process image")
		}

		return c.Send(buf)
	}

	if c.Query("saturation") != "" {

		saturation, err := strconv.ParseFloat(c.Query("saturation"), 64)
		if err != nil {
			slog.Error("saturation should be an int")
			return c.Status(400).SendString("saturation should be an int")
		}

		if saturation > 0 {
			buf, err := processors.Saturation(image, saturation, imageType)
			if err != nil {
				return c.Status(400).SendString("Failed to process image")
			}

			return c.Send(buf)
		}
	}

	if c.Query("fit-w") != "" && c.Query("fit-h") != "" {
		w, err := strconv.Atoi(c.Query("fit-w"))
		if err != nil {
			slog.Error("fit width should be an int")
			return c.Status(400).SendString("width should be an int")
		}
		h, err := strconv.Atoi(c.Query("fit-h"))
		if err != nil {
			slog.Error("fit height should be an int")
			return c.Status(400).SendString("height should be an int")
		}

		buf, err := processors.Fit(image, w, h, imaging.Lanczos, imageType)
		if err != nil {
			return c.Status(400).SendString("Failed to process image")
		}

		return c.Send(buf)
	}

	return c.Send(buf.Bytes())
}
