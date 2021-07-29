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
	"github.com/luiskeys/ebrain/frame"
	"github.com/luiskeys/ebrain/video"
)

func main() {
	fmt.Println("eBrain main process started...")

	gl.StartDriver(mainWindow)
}

func mainWindow(driver gxui.Driver) {
	var ppImage image.Image
	// Create process pipeline
	// Channels
	read := make(chan image.Image)
	pp1 := make(chan image.Image)
	pp2 := make(chan image.Image)
	// Go routines
	go video.Read(read)
	go frame.Preprocess1(pp1, read)
	go frame.Preprocess2(pp2, pp1)
	ppImage = <-pp1

	window, theme, width, height := createWindow(ppImage, driver)

	pause := time.Millisecond * 2
	var timer *time.Timer
	var t0 int64 = time.Now().UnixNano() / int64(time.Millisecond)
	var fps int64
	counter := 0

	timer = time.AfterFunc(pause, func() {
		driver.Call(func() {
			ppImage = <-pp2
			fps, t0 = getFPS(t0)

			updateWindow(window, ppImage, theme, driver, width, height, fps, counter)
			if counter == 10 {
				counter = 0
			}

			counter++

			timer.Reset(pause)
		})
	})
}

func createWindow(ppImage image.Image, driver gxui.Driver) (gxui.Window, gxui.Theme, int, int) {
	// Create window with first frame
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
	return window, theme, width, height
}

func updateWindow(window gxui.Window, ppImage image.Image, theme gxui.Theme, driver gxui.Driver, width int, height int, maxFPS int64, counter int) {

	if counter == 10 {
		window.SetTitle("Frames preview (" + strconv.Itoa(width) + " x " + strconv.Itoa(height) + ") - FPS: " + strconv.Itoa(int(maxFPS)))
	}

	if ppImage != nil {
		img := theme.CreateImage()
		texture := driver.CreateTexture(ppImage, 1.0)
		img.SetTexture(texture)
		window.RemoveAll()
		window.AddChild(img)
	}
}

func getFPS(t0 int64) (int64, int64) {
	t1 := time.Now().UnixNano() / int64(time.Millisecond)
	fps := 1000 / (t1 - t0)
	return fps, t1
}
