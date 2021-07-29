//Frames preprocess entry point
package frame

import (
	"image"
)

func Preprocess2(out chan<- image.Image, in <-chan image.Image) {
	//q := utils.GetPreprocessParams().HorGradientQuantum

	for frame := range in {
		if frame == nil {
			out <- nil
			close(out)
			break
		}

		if frame.Bounds().Max.X < 10 || frame.Bounds().Max.Y < 10 {
			out <- frame
		}

		bufferFrame := Bixels(frame, 10)

		out <- bufferFrame
	}
}

/*
var bufferFrame image.Image
pp := utils.GetPreprocessParams()
drawFrame, bounds := shapes.ImageToDrawImage(bufferFrame)
c := color.RGBA{0, 0, 255, 255}
shapes.DrawGrid(drawFrame, bounds, pp.CellSize, c)
*/
