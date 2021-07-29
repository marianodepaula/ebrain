//Fast conversion to gray for a frame
package frame

import (
	"image"
	"image/color"
)

func Gray(frame image.Image) image.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y

	var r float32 = 0
	var g float32 = 0
	var b float32 = 0

	grayFrame := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			R, G, B, _ := frame.At(x, y).RGBA()
			r = float32(R / 257)
			g = float32(G / 257)
			b = float32(B / 257)
			l := uint8(0.6*r + 0.3*g + 0.1*b)

			grayFrame.Set(x, y, color.Gray{l})
		}
	}

	return grayFrame
}
