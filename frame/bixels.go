package frame

import (
	"image"
	"image/color"
)

func Bixels(frame image.Image, size int) image.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y

	bxFrame := image.NewRGBA(image.Rect(0, 0, width, height))

	semiSize := size / 2
	var ra, ga, ba uint32 = 0, 0, 0
	var cellPixCount uint32 = uint32(size * size)

	for y := semiSize; y < height-semiSize; y += size {
		for x := semiSize; x < width-semiSize; x += size {

			ra = 0
			ga = 0
			ba = 0

			for cy := y - semiSize; cy < y+semiSize; cy++ {
				for cx := x - semiSize; cx < x+semiSize; cx++ {
					R, G, B, _ := frame.At(cx, cy).RGBA()
					ra += R
					ga += G
					ba += B
				}
			}

			ra /= cellPixCount
			ga /= cellPixCount
			ba /= cellPixCount

			ra /= 256
			ga /= 256
			ba /= 256

			for cy := y - semiSize; cy < y+semiSize; cy++ {
				for cx := x - semiSize; cx < x+semiSize; cx++ {
					bxFrame.Set(cx, cy, color.RGBA{uint8(ra), uint8(ga), uint8(ba), 255})
				}
			}

		}
	}

	return bxFrame
}
