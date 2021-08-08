//Calculates the histograms and related functions
package frame

import (
	"image"
	"image/color"
)

func HorizHistoLevels(frame image.Image, col color.Color, vertPerc int, histoThresPerc int) []int {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y
	levels := make([]int, width)

	heightScan := height * (100 - vertPerc) / 100
	histoThres := height * histoThresPerc / 100

	counter := 0

	for x := 0; x < width; x++ {
		counter = 0
		for y := height - 1; y >= heightScan; y-- {
			r, g, b, _ := frame.At(x, y).RGBA()
			if r > 0 || g > 0 || b > 0 {
				counter++
			}
		}

		if counter > histoThres {
			levels[x] = counter
		}
	}

	return levels
}

func HorizHisto(frame image.Image, col color.Color, vertPerc int, histoThresPerc int) (image.Image, []int) {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y
	levels := make([]int, width)

	heightScan := height * (100 - vertPerc) / 100
	histoThres := height * histoThresPerc / 100

	histoFrame := image.NewRGBA(image.Rect(0, 0, width, height))

	counter := 0

	for x := 0; x < width; x++ {
		counter = 0
		for y := height - 1; y >= heightScan; y-- {
			r, g, b, _ := frame.At(x, y).RGBA()
			if r > 0 || g > 0 || b > 0 {
				counter++
			}
		}

		if counter > histoThres {
			levels[x] = counter
			for y := height - counter; y < height; y++ {
				histoFrame.Set(x, y, col)
			}
		}
	}

	return histoFrame, levels
}

func VertHisto(frame image.Image, col color.Color, vertPerc int, histoThres int) image.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y

	heightScan := height * (100 - vertPerc) / 100

	histoFrame := image.NewRGBA(image.Rect(0, 0, width, height))

	counter := 0

	for y := height - 1; y >= heightScan; y-- {
		counter = 0
		for x := 0; x < width; x++ {

			r, g, b, _ := frame.At(x, y).RGBA()
			if r > 0 || g > 0 || b > 0 {
				counter++
			}
		}

		if counter > histoThres {
			for x := 0; x < counter; x++ {
				histoFrame.Set(x, y, col)
			}
		}
	}

	return histoFrame
}
