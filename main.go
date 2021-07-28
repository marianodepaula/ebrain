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

	pause := time.Millisecond * 2
	var timer *time.Timer
	var t0 int64 = time.Now().UnixNano() / int64(time.Millisecond)
	var t1 int64
	var fps int64

	timer = time.AfterFunc(pause, func() {
		driver.Call(func() {
			ppImage = <-pp2
			t1 = time.Now().UnixNano() / int64(time.Millisecond)
			fps = 1000 / (t1 - t0)
			t0 = t1
			window.SetTitle("Frames preview (" + strconv.Itoa(width) + " x " + strconv.Itoa(height) + ") - FPS: " + strconv.Itoa(int(fps)))

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
