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
		write(os.Stdout, res)
	}
}
