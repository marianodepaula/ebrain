//Draw a grid with square cells
package shapes

import (
	"image"
	"image/color"
	"image/draw"
)

func DrawGrid(img draw.Image, bounds image.Rectangle, cellSize int, col color.Color) {
	width := bounds.Max.X
	height := bounds.Max.Y

	for y := 0; y < height; y += cellSize {
		if y < height {
			DrawLine(img, 0, y, width-1, y, col)
		}
	}

	for x := 0; x < width; x += cellSize {
		if x < width {
			DrawLine(img, x, 0, x, height-1, col)
		}
	}
}
