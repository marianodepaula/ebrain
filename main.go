package main

import (
	"fmt"
	"os"

	"github.com/luiskeys/ebrain/video"
)

func main() {
	fmt.Println("eBrain main process started...")
	fmt.Println(len(os.Args), os.Args)

	video.Read()
}
