// Provides all the params for the platform in a single place
package utils

type PreprocessParams struct {
	CropLeft          int
	CropTop           int
	CropRight         int
	CropBottom        int
	ReducedFrameWidth int
}

//Number of frames to skip per iteration
const SkipFramesAmount int = 2

func GetPreprocessParams() PreprocessParams {
	//Set preprocess params here
	pp := PreprocessParams{}
	pp.ReducedFrameWidth = 600
	pp.CropLeft = 250
	pp.CropTop = 380
	pp.CropRight = 180
	pp.CropBottom = 50

	return pp
}
