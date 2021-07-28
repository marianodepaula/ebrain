// Provides all the params for the platform in a single place
package utils

type PreprocessParams struct {
	CropLeft           int
	CropTop            int
	CropRight          int
	CropBottom         int
	ReducedFrameWidth  int
	CellSize           int
	HorGradientQuantum int64
}

//Number of frames to skip per iteration
const SkipFramesAmount int = 0

func GetPreprocessParams() PreprocessParams {
	//Set preprocess params here
	pp := PreprocessParams{}
	pp.ReducedFrameWidth = 600
	pp.CropLeft = 300
	pp.CropTop = 380
	pp.CropRight = 180
	pp.CropBottom = 180
	pp.CellSize = 30
	pp.HorGradientQuantum = 15000

	return pp
}
