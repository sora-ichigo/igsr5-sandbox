package main

import "testing"

// small image
func BenchmarkConvertWithSmallSize(b *testing.B) { convert("./small.png") }

// large image
func BenchmarkConvertWithLargeSize(b *testing.B) { convert("./large.png") }
