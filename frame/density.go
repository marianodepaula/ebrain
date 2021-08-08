//Scan areas with more density based on a black and color image
package frame

import (
	"image"
	"image/color"
)

func ScanByDensity(frame image.Image, col color.Color, scanY0Perc int, scanY1Perc int, thersPerc int) image.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y
	winWidh := width * 20 / 100
	//Area to be scanned
	scanY0 := height * scanY0Perc / 100
	scanY1 := height * scanY1Perc / 100
	scanHeight := scanY1 - scanY0

	scanFrame := image.NewRGBA(image.Rect(0, 0, width, height))
	densArea := scanHeight * winWidh / 100

	dens := 0
	for w := 0; w < width; w += winWidh {
		dens = 0
		for y := scanY0; y < scanY1; y++ {
			for x := w; x < w+winWidh; x++ {
				r, g, b, _ := frame.At(x, y).RGBA()
				if r > 0 || g > 0 || b > 0 {
					dens++
				}
			}
		}

		dens /= densArea

		if dens >= thersPerc {
			for y := 0; y < height; y++ {
				for x := w; x < w+winWidh; x++ {
					scanFrame.Set(x, y, col)
				}
				//scanFrame.Set(w, y, col)
				//scanFrame.Set(w+width, y, col)
			}

		}

	}

	return scanFrame
}
