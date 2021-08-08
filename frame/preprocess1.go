//Frames preprocess entry point
package frame

import (
	"image"

	"github.com/luiskeys/ebrain/utils"
)

func Preprocess1(out chan<- image.Image, in <-chan image.Image) {
	bpp := utils.GetPreprocessParams().ByPassPreprocess1
	pp := utils.GetPreprocessParams()

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
				bufferFrame := Crop(frame, pp.CropLeft, pp.CropTop, pp.CropRight, pp.CropBottom)
				//bufferFrame = Blur(bufferFrame)
				out <- bufferFrame
			}
		}
	}
}
