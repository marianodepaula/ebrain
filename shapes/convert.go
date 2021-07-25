package shapes

import (
	"image"
	"image/draw"
)

func ImageToDrawImage(source image.Image) (draw.Image, image.Rectangle) {
	bounds := source.Bounds()
	bufferDraw := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(bufferDraw, bufferDraw.Bounds(), source, bounds.Min, draw.Src)
	return bufferDraw, bounds
}
