package main

import (
	"fmt"
	"os"

	_ "net/http/pprof"

	"github.com/h2non/bimg"
)

func main() {
	// Resize("img/src_w_13000.png", 2560)
	Convert("img/src_w_13000.png", bimg.WEBP)
}

func Resize(filename string, w int) {
	buffer, err := bimg.Read(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	img := bimg.NewImage(buffer)

	_, err = img.Process(bimg.Options{Width: w, Enlarge: true})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	bimg.Write("dst", img.Image())
}

func Convert(filename string, imgType bimg.ImageType) {
	buffer, err := bimg.Read(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	img := bimg.NewImage(buffer)

	_, err = img.Process(bimg.Options{Type: imgType, Width: 2560})
	// _, err = img.Process(bimg.Options{Width: w, Enlarge: true})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("dst", img.Image())
}

//
// func extract(filename string) {
// 	buffer, err := bimg.Read(filename)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}
//
// 	bimg.NewImage(buffer).Extract(500, 500, 30, 30)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}
// }
//
// func crop(filename string) {
// 	buffer, err := bimg.Read(filename)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}
// 	options := bimg.Options{
// 		Width:   500,
// 		Height:  500,
// 		Crop:    true,
// 		Gravity: bimg.GravityCentre,
// 	}
// 	bimg.NewImage(buffer).Process(options)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}
// }
