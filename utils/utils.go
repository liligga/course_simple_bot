package utils

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
)

func OpenImage(imagePath string) image.Image {

	imgFile, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Error while opening image", err)
	}

	// fmt.Println("Image opened", imgFile)

	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("Error while decoding image", err)
	}

	// fmt.Println("Image decoded", img)

	return img
}
