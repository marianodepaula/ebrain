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

const REDUCED_WIDTH int = 300

func Reduce(frame image.Image) image.Image {

	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y

	xinc := width / REDUCED_WIDTH
	reducedHeight := height * REDUCED_WIDTH / width
	yinc := height / reducedHeight

	//fmt.Println("xinc %d", xinc)
	//fmt.Println("yinc %d", yinc)

	xlim := 0
	ylim := 0
	xreduced := 0
	yreduced := 0
	widthdec := width - 1
	heightdec := height - 1

	reducedFrame := image.NewGray(image.Rect(0, 0, REDUCED_WIDTH, reducedHeight))

	var r uint8 = 0
	//var g uint8 = 0
	//var b uint8 = 0

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

			R, _, _, _ := frame.At(xlim, ylim).RGBA()
			r = uint8(R / 257)
			//g = uint8(G / 257)
			//b = uint8(B / 257)

			reducedFrame.Set(xreduced, yreduced, color.Gray{r})

			xreduced++
		}

		yreduced++
	}

	return reducedFrame
}
