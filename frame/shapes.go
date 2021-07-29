package frame

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/luiskeys/ebrain/shapes"
)

func FindVertLines(frame image.Image, col color.Color) draw.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y
	var r, g, b uint32
	var t uint32
	var y1, y2 int

	bufferDraw := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		y1, y2 = 0, 0
		for y := 0; y < height; y++ {
			r, g, b, _ = frame.At(x, y).RGBA()
			t = r + g + b
			if t > 0 && y1 == 0 {
				y1 = y
			}
			if t == 0 && y1 > 0 {
				y2 = y
				if y2-y1 > 20 {
					shapes.DrawLine(bufferDraw, x, y1, x, y2, col)
				}
				y1 = 0
			}
		}
	}
	return bufferDraw
}

func FindHorizLines(frame image.Image, col color.Color) draw.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y
	var r, g, b uint32
	var t uint32
	var x1, x2 int

	bufferDraw := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		x1, x2 = 0, 0
		for x := 0; x < width; x++ {
			r, g, b, _ = frame.At(x, y).RGBA()
			t = r + g + b
			if t > 0 && x1 == 0 {
				x1 = x
			}
			if t == 0 && x1 > 0 {
				x2 = x
				if x2-x1 > 80 {
					shapes.DrawLine(bufferDraw, x1, y, x2, y, col)
				}
				x1 = 0
			}
		}
	}
	return bufferDraw
}
