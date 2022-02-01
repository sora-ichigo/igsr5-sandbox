package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"

	"github.com/h2non/bimg"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

}

// func convert(b *testing.B, filename string) {
// 	buffer, err := bimg.Read(filename)
// 	if err != nil {
// 		b.Fatalf("failed to Read file. err: %v", err)
// 	}
//
// 	newImage, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
// 	if err != nil {
// 		b.Fatalf("failed to convert. err: %v", err)
// 	}
//
// 	if bimg.NewImage(newImage).Type() == "png" {
// 		b.Fatalf("bad convert.\nwant: %s\ngot: %s", "webp", bimg.NewImage(newImage).Type())
// 	}
// }

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
