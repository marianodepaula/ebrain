//Frames preprocess entry point
package frame

import "image"

type PreprocessParams struct {
	ReducedFrame int
}

func Preprocess(frame image.Image) image.Image {

	preprocessParams := PreprocessParams{}
	preprocessParams.ReducedFrame = 400

	reducedFrame := Reduce(frame, preprocessParams.ReducedFrame)
	grayFrame := ToGray(reducedFrame)
	return grayFrame
}
