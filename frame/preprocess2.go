//Frames preprocess entry point
package frame

import (
	"image"
)

func Preprocess2(out chan<- image.Image, in <-chan image.Image) {
	var bufferFrame image.Image
	for frame := range in {
		if frame == nil {
			out <- nil
			close(out)
			break
		}

		if frame.Bounds().Max.X < 10 || frame.Bounds().Max.Y < 10 {
			out <- frame
		}

		bufferFrame = GetChannels(frame, true, true, false)
		out <- bufferFrame
	}
}
