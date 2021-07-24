//Fast conversion to gray for a frame
package frame

import (
	"image"
	"image/color"
)

func Blur(frame image.Image) image.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y

	blurFrame := image.NewGray(image.Rect(0, 0, width, height))

	var l color.Color
	var m uint32
	var w uint32
	var n uint32
	var e uint32
	var s uint32

	widthdec := width - 1
	heightdec := height - 1

	for y := 0; y < height; y++ {
		l = frame.At(0, y)
		blurFrame.Set(0, y, l)

		l = frame.At(widthdec, y)
		blurFrame.Set(widthdec, y, l)
	}

	for x := 0; x < width; x++ {
		l = frame.At(x, 0)
		blurFrame.Set(x, 0, l)

		l = frame.At(x, heightdec)
		blurFrame.Set(x, heightdec, l)
	}

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {

			l = frame.At(x-1, y)
			m, _, _, _ = l.RGBA()
			_ = m

			l = frame.At(x-1, y)
			w, _, _, _ = l.RGBA()
			l = frame.At(x, y+1)
			n, _, _, _ = l.RGBA()
			l = frame.At(x+1, y)
			e, _, _, _ = l.RGBA()
			l = frame.At(x, y-1)
			s, _, _, _ = l.RGBA()

			m = (w + n + e + s) / 1028

			blurFrame.Set(x, y, color.Gray{uint8(m)})

		}
	}

	return blurFrame
}
