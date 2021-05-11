package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"

	_ "github.com/gonutz/bmp"
)

func main() {
	img, _, err := image.Decode(os.Stdin)
	if err == nil {
		png.Encode(os.Stdout, img)
	}
}
