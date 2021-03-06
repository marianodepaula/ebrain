// Video reader package based on ffmpeg
package video

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/luiskeys/ebrain/utils"
)

type VideoPaths struct {
	inputVideoFolder     string
	sourceFramesFolder   string
	inputVideoFile       string
	inputVideoFileName   string
	inputFileName        string
	sourceFramesFileName string
	firstFrameFile       string
}

func Read(out chan<- image.Image) {

	videoPaths := loadVideoPaths()

	if !utils.FileExists(videoPaths.firstFrameFile, videoPaths.sourceFramesFolder) {
		cmd := exec.Command("ffmpeg", "-i", videoPaths.inputFileName, "-frames:v", "250", videoPaths.sourceFramesFolder+videoPaths.sourceFramesFileName)

		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	framesFiles := make([]string, 0)
	framesFiles = utils.GetFiles(videoPaths.sourceFramesFolder, framesFiles)

	skipIndex := 0
	maxFrames := utils.GetPreprocessParams().MaxFramesToPlay

	for i, f := range framesFiles {
		if skipIndex >= utils.SkipFramesAmount && i < maxFrames {
			out <- loadFrame(videoPaths, f)
			skipIndex = 0
		}

		skipIndex++
	}

	out <- nil
	close(out)
}

func loadFrame(videoPaths VideoPaths, fileName string) image.Image {
	f, err := os.Open(videoPaths.sourceFramesFolder + fileName)
	if err != nil {
		panic("Frame file not found")
	}

	imageData, err := png.Decode(f)
	defer f.Close()
	if err != nil {
		panic("Source frame not found")
	}

	return imageData
}

func loadVideoPaths() VideoPaths {

	videoPahts := VideoPaths{
		utils.GetPath(utils.InputVideo),
		utils.GetPath(utils.SourceFrames),
		"utils.GetFirstFile(inputVideoFolder)",
		"",
		"",
		"",
		"",
	}

	videoPahts.inputVideoFile = utils.GetFirstFile(videoPahts.inputVideoFolder)
	videoPahts.inputVideoFileName = strings.TrimSuffix(videoPahts.inputVideoFile, filepath.Ext(videoPahts.inputVideoFile))
	videoPahts.inputFileName = videoPahts.inputVideoFolder + videoPahts.inputVideoFile
	videoPahts.sourceFramesFileName = videoPahts.inputVideoFileName + "-%04d.png"
	videoPahts.firstFrameFile = videoPahts.inputVideoFileName + "-0001.png"

	return videoPahts
}
