//Merge 2 images with a mix percentage relation
package frame

import (
	"image"
	"image/color"
)

func Overlay(frame1 image.Image, frame2 image.Image, mixperc uint32) image.Image {
	width := frame1.Bounds().Max.X
	height := frame1.Bounds().Max.Y
	width2 := frame2.Bounds().Max.X
	height2 := frame2.Bounds().Max.Y

	c1 := mixperc
	c2 := 100 - mixperc

	if width != width2 || height != height2 {
		return frame1
	}

	var r, g, b uint32 = 0, 0, 0
	var r1, g1, b1 uint32 = 0, 0, 0
	var r2, g2, b2 uint32 = 0, 0, 0

	mergedFrame := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			r1, g1, b1, _ = frame1.At(x, y).RGBA()
			r2, g2, b2, _ = frame2.At(x, y).RGBA()

			r = r1*c1 + r2*c2
			g = g1*c1 + g2*c2
			b = b1*c1 + b2*c2

			r /= 25600
			g /= 25600
			b /= 25600

			mergedFrame.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
	}

	return mergedFrame
}

func Merge(frame1 image.Image, frame2 image.Image, col color.Color) image.Image {
	width := frame1.Bounds().Max.X
	height := frame1.Bounds().Max.Y
	width2 := frame2.Bounds().Max.X
	height2 := frame2.Bounds().Max.Y

	if width != width2 || height != height2 {
		return frame1
	}

	var r, g, b uint32 = 0, 0, 0
	var r1, g1, b1 uint32 = 0, 0, 0
	var r2, g2, b2 uint32 = 0, 0, 0

	mergedFrame := image.NewRGBA(image.Rect(0, 0, width, height))

	var dyncol color.Color = color.Black

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			r1, g1, b1, _ = frame1.At(x, y).RGBA()
			r2, g2, b2, _ = frame2.At(x, y).RGBA()

			r = r1 + r2
			g = g1 + g2
			b = b1 + b2

			dyncol = color.Black
			if r > 0 || g > 0 || b > 0 {
				dyncol = col
			}

			mergedFrame.Set(x, y, dyncol)
		}
	}

	return mergedFrame
}
