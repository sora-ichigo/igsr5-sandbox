package main

import (
	"fmt"
	"os"

	_ "net/http/pprof"

	"github.com/h2non/bimg"
)

func main() {
	Resize("img/src_w_6000.png", 3000, 600)
	// Convert("img/src.png", bimg.WEBP)
}

func Resize(filename string, w int, h int) {
	buffer, err := bimg.Read(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	img := bimg.NewImage(buffer)

	_, err = img.Process(bimg.Options{Width: w, Height: h, Enlarge: true, Crop: true})
	// 	_, err = img.Process(bimg.Options{Width: w, Enlarge: true})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// bimg.Write("out/dst.png", img.Image())
}

func Convert(filename string, imgType bimg.ImageType) {
	buffer, err := bimg.Read(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	img := bimg.NewImage(buffer)

	_, err = img.Process(bimg.Options{Type: imgType})
	// _, err = img.Process(bimg.Options{Width: w, Enlarge: true})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("out/webp.png", img.Image())
}

func extract(filename string) {
	buffer, err := bimg.Read(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.NewImage(buffer).Extract(500, 500, 30, 30)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func crop(filename string) {
	buffer, err := bimg.Read(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	options := bimg.Options{
		Width:   500,
		Height:  500,
		Crop:    true,
		Gravity: bimg.GravityCentre,
	}
	bimg.NewImage(buffer).Process(options)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
