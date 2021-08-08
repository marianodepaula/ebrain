// Provides all the params for the platform in a single place
package utils

type PreprocessParams struct {
	MaxFramesToPlay   int
	ByPassPreprocess1 bool
	ByPassPreprocess2 bool
	ByPassPreprocess3 bool
	CropLeft          int
	CropTop           int
	CropRight         int
	CropBottom        int
	ReducedFrameWidth int
	CellSize          int
	HorizGradQuantum  int64
	ScanY0Perc        int
	ScanY1Perc        int
	ThersPerc         int
}

//Number of frames to skip per iteration
const SkipFramesAmount int = 0

func GetPreprocessParams() PreprocessParams {
	//Set preprocess params here
	pp := PreprocessParams{}
	pp.MaxFramesToPlay = 50
	pp.ByPassPreprocess1 = false
	pp.ByPassPreprocess2 = false
	pp.ByPassPreprocess3 = false
	pp.CropLeft = 300
	pp.CropTop = 380
	pp.CropRight = 180
	pp.CropBottom = 180
	pp.CellSize = 30
	pp.HorizGradQuantum = 10000
	pp.ScanY0Perc = 30
	pp.ScanY1Perc = 70
	pp.ThersPerc = 20

	return pp
}
