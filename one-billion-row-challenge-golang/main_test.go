package main

import (
	"os"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	for b.Loop() {
		f, err := os.Open("measurements_7.txt")
		if err != nil {
			panic(err)
		}

		res := run(f)

		f.Close()
		write(os.Stdout, res)
	}
}
