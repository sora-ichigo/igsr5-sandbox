package main

import "testing"

func BenchmarkResize(b *testing.B) {
	tests := []struct {
		name     string
		filename string
		width    int
	}{
		{name: "width 1280px to 300px", filename: "img/src_w_1280.jpeg", width: 300},
		{name: "width 320px to 300px", filename: "img/src_w_320.jpeg", width: 300},
		{name: "width 1280px to 320px", filename: "img/src_w_1280.jpeg", width: 320},
		{name: "width 320px to 320px", filename: "img/src_w_320.jpeg", width: 320},
		{name: "width 2560px to 1500px", filename: "img/src_w_2560.jpeg", width: 1500},
		{name: "width 1920px to 1500px", filename: "img/src_w_1920.jpeg", width: 1500},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Resize(tt.filename, tt.width)
			}
		})
	}
}

// func BenchmarkConvert(b *testing.B) {
// 	tests := []string{"img/src.png", "img/src.webp"}
//
// 	for _, t := range tests {
// 		b.Run(t, func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				Convert(t, bimg.WEBP)
// 			}
// 		})
// 	}
// }

// target function. https://github.com/h2non/bimg#convert
// func convert(b *testing.B, filename string) {
// 	b.Helper()
//
// 	buffer, err := bimg.Read(filename)
// 	if err != nil {
// 		b.Fatalf("failed to read file. err: %v", err)
// 	}
//
// 	src := bimg.NewImage(buffer)
//
// 	src.Convert(bimg.WEBP)
// }
//
// func resizeAndConvert(b *testing.B, filename string) {
// 	b.Helper()
//
// 	buffer, err := bimg.Read(filename)
// 	if err != nil {
// 		b.Fatalf("failed to read file. err: %v", err)
// 	}
//
// 	src := bimg.NewImage(buffer)
//
// 	src.Resize(900, 500)
// 	src.Convert(bimg.WEBP)
// }
//
// func resize(b *testing.B, filename string) {
// 	b.Helper()
//
// 	buffer, err := bimg.Read(filename)
// 	if err != nil {
// 		b.Fatalf("failed to read file. err: %v", err)
// 	}
//
// 	src := bimg.NewImage(buffer)
//
// 	src.Resize(6000, 4000)
// }
//
// // benchmark of convert()
//
// // 3MB, 900x500px
// func BenchmarkConvertWithSmallSize(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		resizeAndConvert(b, "./small.png")
// 	}
// }
//
// // 3MB, 6000x4000px
// func BenchmarkConvertWithLargeSize(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		resizeAndConvert(b, "./large.png")
// 	}
// }
//
// // 3MB, 900x500px
// func BenchmarkResizeWithSmallSize(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		resize(b, "./small.png")
// 	}
// }
//
// // 3MB, 6000x4000px
// func BenchmarkResizeWithLargeSize(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		resize(b, "./large.png")
// 	}
// }
//
// // Extract
// // func BenchmarkExtract(b *testing.B) {
// // 	for i := 0; i < b.N; i++ {
// // 		extract("./large.png")
// // 	}
// // }
// //
// // Crop
// // func BenchmarkCrop(b *testing.B) {
// // 	for i := 0; i < b.N; i++ {
// // 		crop("./large.png")
// // 	}
// // }
