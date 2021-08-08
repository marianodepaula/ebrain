//Frames preprocess 2
package frame

import (
	"image"
	"image/color"

	"github.com/luiskeys/ebrain/utils"
)

func Preprocess2(out chan<- image.Image, in <-chan image.Image) {
	bpp := utils.GetPreprocessParams().ByPassPreprocess2
	q := utils.GetPreprocessParams().HorizGradQuantum
	col := color.RGBA{0, 255, 0, 255}

	for frame := range in {
		if frame == nil {
			out <- nil
			close(out)
			return
		}

		if frame.Bounds().Max.X < 10 || frame.Bounds().Max.Y < 10 {
			//Frame is too small, then bypass
			out <- frame
		} else {
			if bpp {
				//Bypass, no process
				out <- frame
			} else {
				//Proces here
				horizGradFrame := HorizGrad(frame, q)
				vertGradFrame := VertGrad(frame, q)
				gradFrame := Merge(horizGradFrame, vertGradFrame, col)
				out <- gradFrame
			}

		}
	}
}
