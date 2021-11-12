package main

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
)

//go:embed img/*
var testImages embed.FS

// ファイルサイズ自体は小さいが image.Image に展開するとメモリをとても食う画像処理の負荷実験
func main() {
	// original png
	srcPng := decodeFile("img/large-black-original.png")
	img, _, err := image.Decode(bytes.NewReader(srcPng))
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	fmt.Printf("srcPng Bounds %v", img.Bounds())
	fmt.Println("")

	// 5000px png
	srcPng5000 := decodeFile("img/large-black-5000.png")
	img, _, err = image.Decode(bytes.NewReader(srcPng5000))
	if err != nil {
		log.Fatalf("5000 err: %v", err)
	}
	fmt.Printf("srcPng5000 Bounds %v", img.Bounds())
	fmt.Println("")

	// 10000px png
	srcPng10000 := decodeFile("img/large-black-10000.png")
	img, _, err = image.Decode(bytes.NewReader(srcPng10000))
	if err != nil {
		log.Fatalf("10000 err: %v", err)
	}
	fmt.Printf("srcPng10000 Bounds %v", img.Bounds())
	fmt.Println("")

	// 20000px png
	srcPng20000 := decodeFile("img/large-black-20000.png")
	img, _, err = image.Decode(bytes.NewReader(srcPng20000))
	if err != nil {
		log.Fatalf("20000 err: %v", err)
	}
	fmt.Printf("srcPng20000 Bounds: %v", img.Bounds())
}

// decode from embed
func decodeFile(filename string) []byte {
	b, err := testImages.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file: %s", filename)
	}

	return b
}
