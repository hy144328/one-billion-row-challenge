package main

import (
	"flag"
	"os"
	"testing"
)

var inPath = flag.String("in", "measurements.txt", "Input file with measurements.")

func BenchmarkRun(b *testing.B) {
	f, err := os.Open(*inPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for b.Loop() {
		f.Seek(0, 0)
		res := run(f)
		writeFloat(os.Stdout, res)
	}
}

func BenchmarkRun1(b *testing.B) {
	f, err := os.Open(*inPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for b.Loop() {
		f.Seek(0, 0)
		res := run1(f)
		writeFloat(os.Stdout, res)
	}
}

func BenchmarkRun2(b *testing.B) {
	f, err := os.Open(*inPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for b.Loop() {
		f.Seek(0, 0)
		res := run2(f)
		writeInt(os.Stdout, res)
	}
}
