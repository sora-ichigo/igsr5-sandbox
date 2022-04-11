package main

import (
	"fmt"
	"os"

	_ "net/http/pprof"

	"github.com/h2non/bimg"
)

func main() {
	Resize("img/src_w_13000.jpeg", 2560)
	Resize("img/src_w_13000.jpeg", 1920)
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

	// dst := fmt.Sprintf("img/src_w_%d.%s", w, img.Type())
	// bimg.Write(dst, img.Image())
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
