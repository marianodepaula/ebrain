package main

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/themes/dark"
	"github.com/luiskeys/ebrain/video"
)

func main() {
	fmt.Println("eBrain main process started...")

	gl.StartDriver(mainWindow)
}

func mainWindow(driver gxui.Driver) {
	var ppImage image.Image = video.Read()
	width, height := ppImage.Bounds().Max.X, ppImage.Bounds().Max.Y
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(m, ppImage.Bounds(), ppImage, image.Point{}, draw.Src)

	theme := dark.CreateTheme(driver)
	img := theme.CreateImage()
	window := theme.CreateWindow(width, height, "Image viewer")
	texture := driver.CreateTexture(m, 1.0)
	img.SetTexture(texture)
	window.AddChild(img)
	window.OnClose(driver.Terminate)
}
