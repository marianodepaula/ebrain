// Provides path management
package utils

import (
	"io/ioutil"
)

const BasePath = ""
const InputVideo = "input_video"
const SourceFrames = "source_frames"

func GetPath(path string) string {
	switch path {
	case InputVideo:
		return BasePath + "inputvideo/"
	case SourceFrames:
		return BasePath + "sourceframes/"
	default:
		return BasePath
	}
}

func FileExists(fileName string, path string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic("Fali to read files in " + fileName)
	}

	for _, f := range files {
		return f.Name() == fileName
	}

	return false
}

func GetFiles(path string, filesList []string) []string {

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic("Fali to read files in " + path)
	}

	for _, f := range files {
		filesList = append(filesList, f.Name())
	}

	return filesList
}

func GetFirstFile(path string) string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic("Fali to read files in " + path)
	}

	for _, f := range files {
		return f.Name()
	}

	return ""
}
