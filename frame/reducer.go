//Fast downscale a frame
package frame

import (
	"image"
	"image/color"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

func Reduce(frame image.Image, reducedWidth int) image.Image {

	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y

	xinc := width / reducedWidth
	reducedHeight := height * reducedWidth / width
	yinc := height / reducedHeight

	//fmt.Println("xinc %d", xinc)
	//fmt.Println("yinc %d", yinc)

	xlim := 0
	ylim := 0
	xreduced := 0
	yreduced := 0
	widthdec := width - 1
	heightdec := height - 1

	reducedFrame := image.NewRGBA(image.Rect(0, 0, reducedWidth, reducedHeight))

	var r uint8 = 0
	var g uint8 = 0
	var b uint8 = 0

	for y := 0; y < height; y += yinc {
		ylim = y

		if y > heightdec {
			ylim = heightdec
		}

		xreduced = 0
		for x := 0; x < width; x += xinc {
			xlim = x

			if x > widthdec {
				xlim = widthdec
			}

			R, G, B, _ := frame.At(xlim, ylim).RGBA()
			r = uint8(R / 257)
			g = uint8(G / 257)
			b = uint8(B / 257)

			reducedFrame.Set(xreduced, yreduced, color.RGBA{r, g, b, 255})

			xreduced++
		}

		yreduced++
	}

	return reducedFrame
}
