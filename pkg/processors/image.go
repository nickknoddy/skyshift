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
