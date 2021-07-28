//Frames preprocess entry point
package frame

import (
	"image"

	"github.com/luiskeys/ebrain/utils"
)

func Preprocess1(out chan<- image.Image, in <-chan image.Image) {
	//Set preprocess params here
	pp := utils.GetPreprocessParams()
	for frame := range in {
		if frame == nil {
			out <- nil
			close(out)
			break
		}

		if frame.Bounds().Max.X < 10 || frame.Bounds().Max.Y < 10 {
			out <- frame
		}

		bufferFrame := Crop(frame, pp.CropLeft, pp.CropTop, pp.CropRight, pp.CropBottom)
		bufferFrame = Blur(bufferFrame)

		out <- bufferFrame
	}
}
