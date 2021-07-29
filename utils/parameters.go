// Provides all the params for the platform in a single place
package utils

type PreprocessParams struct {
	ByPassPreprocess1 bool
	ByPassPreprocess2 bool
	CropLeft          int
	CropTop           int
	CropRight         int
	CropBottom        int
	ReducedFrameWidth int
	CellSize          int
	HorizGradQuantum  int64
	BixelSize         int
}

//Number of frames to skip per iteration
const SkipFramesAmount int = 0

func GetPreprocessParams() PreprocessParams {
	//Set preprocess params here
	pp := PreprocessParams{}
	pp.ByPassPreprocess1 = false
	pp.ByPassPreprocess2 = false
	pp.ReducedFrameWidth = 600
	pp.CropLeft = 300
	pp.CropTop = 380
	pp.CropRight = 180
	pp.CropBottom = 180
	pp.CellSize = 30
	pp.HorizGradQuantum = 10000
	pp.BixelSize = 40

	return pp
}
