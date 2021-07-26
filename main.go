//Foundation POC for Cognitive Autonomous Machines and AGI
//The main goal is to mix techniques to reduced required calculations in one or more orders
package main

import (
	"fmt"
	"image"
	"image/draw"
	"strconv"
	"time"

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
	var ppImage image.Image
	ppframes := make(chan image.Image)

	go video.Read(ppframes)
	ppImage = <-ppframes

	width, height := ppImage.Bounds().Max.X, ppImage.Bounds().Max.Y
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(m, ppImage.Bounds(), ppImage, image.Point{}, draw.Src)

	theme := dark.CreateTheme(driver)
	img := theme.CreateImage()
	window := theme.CreateWindow(width, height, "Frames preview ("+strconv.Itoa(width)+", "+strconv.Itoa(height)+")")
	texture := driver.CreateTexture(m, 1.0)
	img.SetTexture(texture)
	window.AddChild(img)
	window.OnClose(driver.Terminate)

	pause := time.Millisecond * 5
	var timer *time.Timer
	timer = time.AfterFunc(pause, func() {
		driver.Call(func() {
			ppImage = <-ppframes

			if ppImage != nil {
				img := theme.CreateImage()
				texture := driver.CreateTexture(ppImage, 1.0)
				img.SetTexture(texture)
				window.RemoveAll()
				window.AddChild(img)
			}

			timer.Reset(pause)
		})
	})
}
