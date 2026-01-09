package main

import (
	"flag"
	"os"
	"testing"
)

var measurementsPath = flag.String("measurements", "measurements_7.txt", "Input file with measurements.")

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		f, err := os.Open(*measurementsPath)
		if err != nil {
			panic(err)
		}

		res := run(f)

		f.Close()
		write(os.Stdout, res)
	}
}
