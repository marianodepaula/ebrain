package frame

import (
	"image"
	"image/color"
)

func GetHorGradient(frame image.Image, quantum int64) image.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y
	var w, e uint32
	var r, g, b uint32

	gradientFrame := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			r, g, b, _ = frame.At(x-1, y).RGBA()
			w = r + g + b

			r, g, b, _ = frame.At(x+1, y).RGBA()
			e = r + g + b

			col := getColor(w, e, quantum)

			gradientFrame.Set(x, y, col)
		}
	}

	return gradientFrame
}

func getColor(w uint32, e uint32, quantum int64) color.Color {

	wi := int64(w)
	ei := int64(e)

	if (wi - ei) > quantum {
		return color.RGBA{0, 255, 0, 255}
	}

	if ei-wi > quantum {
		return color.RGBA{0, 255, 0, 255}

	}

	return color.Black
}
