package frame

import (
	"image"
	"image/color"
)

func GetChannels(frame image.Image, red bool, green bool, blue bool) image.Image {
	width := frame.Bounds().Max.X
	height := frame.Bounds().Max.Y

	filteredFrame := image.NewRGBA(image.Rect(0, 0, width, height))

	chFlag := 0
	getChFlag(&chFlag, red, green, blue)

	switch chFlag {
	case 4: // r
		getRed(frame, filteredFrame, width, height, red, green, blue)
	case 2: // g
		getGreen(frame, filteredFrame, width, height, red, green, blue)
	case 1: // b
		getBlue(frame, filteredFrame, width, height, red, green, blue)
	case 6: // r, g
		getRedGreen(frame, filteredFrame, width, height, red, green, blue)
	case 5: // r, b
		getRedGreen(frame, filteredFrame, width, height, red, green, blue)
	case 3: // g, b
		getGreenBlue(frame, filteredFrame, width, height, red, green, blue)
	default: // r, g, b or none
		return frame
	}

	return filteredFrame
}

func getChFlag(chFlag *int, red bool, green bool, blue bool) {

	if red {
		*chFlag += 4
	}

	if green {
		*chFlag += 2
	}

	if blue {
		*chFlag += 1
	}
}

func getRed(frame image.Image, filteredFrame *image.RGBA, width int, height int, bool, green bool, blue bool) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, _, _, _ := frame.At(x, y).RGBA()
			filteredFrame.Set(x, y, color.RGBA{uint8(r), uint8(0), uint8(0), 255})
		}
	}
}

func getGreen(frame image.Image, filteredFrame *image.RGBA, width int, height int, bool, green bool, blue bool) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			_, g, _, _ := frame.At(x, y).RGBA()
			filteredFrame.Set(x, y, color.RGBA{uint8(0), uint8(g), uint8(0), 255})
		}
	}
}

func getBlue(frame image.Image, filteredFrame *image.RGBA, width int, height int, bool, green bool, blue bool) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			_, _, b, _ := frame.At(x, y).RGBA()
			filteredFrame.Set(x, y, color.RGBA{uint8(0), uint8(0), uint8(b), 255})
		}
	}
}

func getRedGreen(frame image.Image, filteredFrame *image.RGBA, width int, height int, bool, green bool, blue bool) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, _, _ := frame.At(x, y).RGBA()
			filteredFrame.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(0), 255})
		}
	}
}

func getRedBlue(frame image.Image, filteredFrame *image.RGBA, width int, height int, bool, green bool, blue bool) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, _, b, _ := frame.At(x, y).RGBA()
			filteredFrame.Set(x, y, color.RGBA{uint8(r), uint8(0), uint8(b), 255})
		}
	}
}

func getGreenBlue(frame image.Image, filteredFrame *image.RGBA, width int, height int, bool, green bool, blue bool) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			_, g, b, _ := frame.At(x, y).RGBA()
			filteredFrame.Set(x, y, color.RGBA{uint8(0), uint8(g), uint8(b), 255})
		}
	}
}
