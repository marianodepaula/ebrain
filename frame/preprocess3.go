//Frames preprocess 3
package frame

import (
	"image"
	"image/color"

	"github.com/luiskeys/ebrain/utils"
)

func Preprocess3(out chan<- image.Image, in <-chan image.Image) {
	bpp := utils.GetPreprocessParams().ByPassPreprocess3
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
				horizLinesFrame := FindHorizLines(frame, col)
				vertLinesFrame := FindVertLines(frame, col)
				linesFrame := Merge(horizLinesFrame, vertLinesFrame, col)
				out <- linesFrame
			}

		}
	}
}
