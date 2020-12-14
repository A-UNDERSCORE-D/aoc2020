package util

import (
	"image"
	"image/png"
	"os"
)

func PNGQuick(targetPath string, img image.Image) {
	f, err := os.Create(targetPath)
	if err != nil {
		panic(err)
	}

	if err := (png.Encode(f, img)); err != nil {
		panic(err)
	}

	f.Close()
}