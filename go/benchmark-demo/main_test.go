package main

import (
	"testing"
)

func BenchmarkItoaByFmt(b *testing.B)      { ItoaByFmt(b.N) }
func BenchmarkItoaByStrconv1(b *testing.B) { ItoaByStrconv1(b.N) }
func BenchmarkItoaByStrconv2(b *testing.B) { ItoaByStrconv2(b.N) }
