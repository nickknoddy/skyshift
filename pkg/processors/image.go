package processors

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"strings"

	"github.com/disintegration/imaging"
)

func InferImageType(filename string) string {

	imageType := strings.Split(filename, ".")
	return imageType[len(imageType)-1]
}

func convertImageFormat(image *image.NRGBA, format string, quality int) ([]byte, error) {

	var buf bytes.Buffer

	if format == "png" {
		if err := png.Encode(&buf, image); err != nil {
			return nil, err
		}
	} else {
		if err := jpeg.Encode(&buf, image, &jpeg.Options{Quality: quality}); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func Resize(image image.Image, w int, h int, filter imaging.ResampleFilter, format string) ([]byte, error) {

	resizedImage := imaging.Resize(image, w, h, filter)
	return convertImageFormat(resizedImage, format, 50)
}

func Sharpen(image image.Image, sharpen float64, format string) ([]byte, error) {

	sharpenImage := imaging.Sharpen(image, sharpen)
	return convertImageFormat(sharpenImage, format, 50)
}

func Blur(image image.Image, blur float64, format string) ([]byte, error) {

	blurImage := imaging.Blur(image, blur)
	return convertImageFormat(blurImage, format, 50)
}

func Brightness(image image.Image, brightness float64, format string) ([]byte, error) {

	brightenedImage := imaging.AdjustBrightness(image, brightness)
	return convertImageFormat(brightenedImage, format, 50)
}

func Contrast(image image.Image, contrast float64, format string) ([]byte, error) {

	contrastImage := imaging.AdjustContrast(image, contrast)
	return convertImageFormat(contrastImage, format, 50)
}

func FlipHorizontal(image image.Image, format string) ([]byte, error) {

	contrastImage := imaging.FlipH(image)
	return convertImageFormat(contrastImage, format, 50)
}

func FlipVertical(image image.Image, format string) ([]byte, error) {

	contrastImage := imaging.FlipV(image)
	return convertImageFormat(contrastImage, format, 50)
}

func GrayScale(image image.Image, format string) ([]byte, error) {

	contrastImage := imaging.Grayscale(image)
	return convertImageFormat(contrastImage, format, 50)
}
