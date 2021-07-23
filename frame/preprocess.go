//Frames preprocess entry point
package frame

import (
	"image"

	"github.com/luiskeys/ebrain/utils"
)

func Preprocess(frame image.Image) image.Image {
	//Set preprocess params here
	pp := utils.GetPreprocessParams()

	croppedFrame := Crop(frame, pp.CropLeft, pp.CropTop, pp.CropRight, pp.CropBottom)
	reducedFrame := Reduce(croppedFrame, pp.ReducedFrameWidth)
	grayFrame := ToGray(reducedFrame)
	return grayFrame
}
