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
	bufferFrame = Reduce(bufferFrame, pp.ReducedFrameWidth)
	bufferFrame = Blur(bufferFrame)
	//bufferFrame = ToGray(bufferFrame)

	//	bufferDraw, bounds := shapes.ImageToDrawImage(bufferFrame)
	//	draw.Draw(bufferDraw, bufferDraw.Bounds(), bufferFrame, bounds.Min, draw.Src)

	//	lineColor := color.RGBA{255, 0, 0, 255}
	//	shapes.DrawLine(bufferDraw, 10, 10, 100, 100, lineColor)
	return bufferFrame
}
