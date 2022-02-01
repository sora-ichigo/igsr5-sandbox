package main

import (
	"testing"

	"github.com/h2non/bimg"
)

// target function. https://github.com/h2non/bimg#convert
func convert(b *testing.B, filename string) {
	b.Helper()

	buffer, err := bimg.Read(filename)
	if err != nil {
		b.Fatalf("failed to read file. err: %v", err)
	}

	src := bimg.NewImage(buffer)

	src.Convert(bimg.WEBP)
}

func resizeAndConvert(b *testing.B, filename string) {
	b.Helper()

	buffer, err := bimg.Read(filename)
	if err != nil {
		b.Fatalf("failed to read file. err: %v", err)
	}

	src := bimg.NewImage(buffer)

	src.Resize(900, 500)
	src.Convert(bimg.WEBP)
}

// benchmark of convert()

// 3MB, 900x500px
func BenchmarkConvertWithSmallSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resizeAndConvert(b, "./small.png")
	}
}

// 3MB, 6000x4000px
func BenchmarkConvertWithLargeSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resizeAndConvert(b, "./large.png")
	}
}

// Extract
func BenchmarkExtract(b *testing.B) {
	extract("./large.png")
}

// Crop
func BenchmarkCrop(b *testing.B) {
	crop("./large.png")
}
