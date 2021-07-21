//Fast frame crop
package frame

import (
	"image"
	"image/color"
)

func Crop(frame image.Image, left int, top int, right int, bottom int) image.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y

	croppedFrame := image.NewRGBA(image.Rect(0, 0, width-left-right, height-top-bottom))

	destx := 0
	desty := 0
	for y := top; y < height-bottom; y++ {
		destx = 0
		for x := left; x < width-right; x++ {
			r, g, b, _ := frame.At(x, y).RGBA()
			croppedFrame.Set(destx, desty, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
			destx++
		}
		desty++
	}

	return croppedFrame
}
