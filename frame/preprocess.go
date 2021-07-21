//Frames preprocess entry point
package frame

import "image"

type PreprocessParams struct {
	CropLeft          int
	CropTop           int
	CropRight         int
	CropBottom        int
	ReducedFrameWidth int
}

func Preprocess(frame image.Image) image.Image {
	//Set preprocess params here
	pp := PreprocessParams{}
	pp.ReducedFrameWidth = 600
	pp.CropLeft = 250
	pp.CropTop = 380
	pp.CropRight = 180
	pp.CropBottom = 50

	croppedFrame := Crop(frame, pp.CropLeft, pp.CropTop, pp.CropRight, pp.CropBottom)
	reducedFrame := Reduce(croppedFrame, pp.ReducedFrameWidth)
	grayFrame := ToGray(reducedFrame)
	return grayFrame
}
