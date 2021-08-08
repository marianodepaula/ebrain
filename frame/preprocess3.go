//Frames preprocess 3
package frame

import (
	"image"
	"image/color"

	"github.com/luiskeys/ebrain/utils"
)

func Preprocess3(out chan<- image.Image, in <-chan image.Image) {
	bpp := utils.GetPreprocessParams().ByPassPreprocess3
	scanY0Perc := utils.GetPreprocessParams().ScanY0Perc
	scanY1Perc := utils.GetPreprocessParams().ScanY1Perc
	thersPerc := utils.GetPreprocessParams().ThersPerc
	col := color.RGBA{0, 0, 255, 255}

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
				scanFrame := ScanByDensity(frame, col, scanY0Perc, scanY1Perc, thersPerc)
				areaFrame := Overlay(frame, scanFrame, 50)
				out <- areaFrame
			}

		}
	}
}
