package frame

import (
	"image"
	"image/color"
)

func HorizGrad(frame image.Image, quantum int64) image.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y
	var w, e uint32
	var r, g, b uint32

	gradientFrame := image.NewRGBA(image.Rect(0, 0, width, height))
	var col color.Color = color.Black
	var colgrad color.Color = color.RGBA{0, 255, 0, 255}
	var wi, ei int64

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			r, g, b, _ = frame.At(x-1, y).RGBA()
			w = r + g + b

			r, g, b, _ = frame.At(x+1, y).RGBA()
			e = r + g + b

			wi = int64(w)
			ei = int64(e)

			col = color.Black

			if (wi - ei) > quantum {
				col = colgrad
			}
			if (ei - wi) > quantum {
				col = colgrad
			}

			gradientFrame.Set(x, y, col)
		}
	}

	return gradientFrame
}
