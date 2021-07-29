//Gets the negative of the frame
package frame

import (
	"image"
	"image/color"
)

func Negative(frame image.Image) image.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y

	negativeFrame := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := frame.At(x, y).RGBA()
			negativeFrame.Set(x, y, color.RGBA{uint8(255 - r), uint8(255 - g), uint8(255 - b), 255})
		}
	}

	return negativeFrame
}
