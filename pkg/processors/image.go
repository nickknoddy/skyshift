package processors

import (
	"strings"
)

func InferImageType(filename string) string {

	imageType := strings.Split(filename, ".")
	return imageType[len(imageType)-1]
}
