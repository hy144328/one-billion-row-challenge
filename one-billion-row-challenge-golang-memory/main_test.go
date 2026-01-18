package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"testing"
)

var inPath = flag.String("in", "measurements.txt", "Input file with measurements.")

func TestRun(t *testing.T) {
	f, err := os.Open("measurements_6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res := run(f)
	var out strings.Builder
	writeInt(&out, res)

	g, err := os.Open("averages_6.txt")
	if err != nil {
		panic(err)
	}
	defer g.Close()

	var sol strings.Builder
	g.WriteTo(&sol)

	if out.String() != sol.String() {
		t.Fail()
	}
}

func BenchmarkRun(b *testing.B) {
	f, err := os.Open(*inPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for b.Loop() {
		f.Seek(0, 0)
		res := run(f)
		writeInt(io.Discard, res)
	}
}
