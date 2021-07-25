//Fast conversion to gray for a frame
package frame

import (
	"image"
	"image/color"
)

func Blur(frame image.Image) image.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y

	blurFrame := image.NewNRGBA(image.Rect(0, 0, width, height))

	var l color.Color
	var m uint32
	var wr, wg, wb uint32
	var nr, ng, nb uint32
	var er, eg, eb uint32
	var sr, sg, sb uint32
	var r, g, b uint32

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
			wr, wg, wb, _ = l.RGBA()
			l = frame.At(x, y+1)
			nr, ng, nb, _ = l.RGBA()
			l = frame.At(x+1, y)
			er, eg, eb, _ = l.RGBA()
			l = frame.At(x, y-1)
			sr, sg, sb, _ = l.RGBA()

			r = (wr + nr + er + sr) / 1028
			g = (wg + ng + eg + sg) / 1028
			b = (wb + nb + eb + sb) / 1028

			blurFrame.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), 255})

		}
	}

	return blurFrame
}
