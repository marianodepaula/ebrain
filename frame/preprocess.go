//Frames preprocess entry point
package frame

import "image"

func Preprocess(frame image.Image) image.Image {
	reducedFrame := Reduce(frame)
	return reducedFrame
}
