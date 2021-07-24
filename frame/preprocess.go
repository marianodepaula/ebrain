//Frames preprocess entry point
package frame

import (
	"image"

	"github.com/luiskeys/ebrain/utils"
)

func Preprocess(frame image.Image) image.Image {
	//Set preprocess params here
	pp := utils.GetPreprocessParams()

	if frame.Bounds().Max.X < 10 || frame.Bounds().Max.Y < 10 {
		return frame
	}

	bufferFrame := Crop(frame, pp.CropLeft, pp.CropTop, pp.CropRight, pp.CropBottom)
	//bufferFrame = Reduce(bufferFrame, pp.ReducedFrameWidth)
	bufferFrame = ToGray(bufferFrame)
	//bufferFrame = Blur(bufferFrame)
	return bufferFrame
}
