package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"testing"
)

var inPath = flag.String("in", "measurements.txt", "Input file with measurements.")

func TestRun0(t *testing.T) {
	f, err := os.Open("measurements_6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res := run0(f)
	var out strings.Builder
	writeFloat(&out, res)

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

func BenchmarkRun0(b *testing.B) {
	f, err := os.Open(*inPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for b.Loop() {
		f.Seek(0, 0)
		res := run0(f)
		writeFloat(io.Discard, res)
	}
}

func TestRun1(t *testing.T) {
	f, err := os.Open("measurements_6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res := run1(f)
	var out strings.Builder
	writeFloat(&out, res)

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

func BenchmarkRun1(b *testing.B) {
	f, err := os.Open(*inPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for b.Loop() {
		f.Seek(0, 0)
		res := run1(f)
		writeFloat(io.Discard, res)
	}
}

func TestRun2(t *testing.T) {
	f, err := os.Open("measurements_6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res := run2(f)
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

func BenchmarkRun2(b *testing.B) {
	f, err := os.Open(*inPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for b.Loop() {
		f.Seek(0, 0)
		res := run2(f)
		writeInt(io.Discard, res)
	}
}

func TestRun3(t *testing.T) {
	f, err := os.Open("measurements_6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res := run3(f)
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

func BenchmarkRun3(b *testing.B) {
	f, err := os.Open(*inPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for b.Loop() {
		f.Seek(0, 0)
		res := run3(f)
		writeInt(io.Discard, res)
	}
}

func TestRun4(t *testing.T) {
	f, err := os.Open("measurements_6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res := run4(f)
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

func BenchmarkRun4(b *testing.B) {
	f, err := os.Open(*inPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for b.Loop() {
		f.Seek(0, 0)
		res := run4(f)
		writeInt(io.Discard, res)
	}
}

func TestRun5(t *testing.T) {
	f, err := os.Open("measurements_6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res := run5(f)
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

func BenchmarkRun5(b *testing.B) {
	f, err := os.Open(*inPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for b.Loop() {
		f.Seek(0, 0)
		res := run5(f)
		writeInt(io.Discard, res)
	}
}

func TestRun6(t *testing.T) {
	f, err := os.Open("measurements_6.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	res := run6(f)
	var out strings.Builder
	writeInt(&out, res.ToMap())

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

func BenchmarkRun6(b *testing.B) {
	f, err := os.Open(*inPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for b.Loop() {
		f.Seek(0, 0)
		res := run6(f)
		writeInt(io.Discard, res.ToMap())
	}
}
