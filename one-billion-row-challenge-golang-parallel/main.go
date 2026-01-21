package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"runtime"
)

const maxBytesPerLine = 128
const maxCities = 10000
const noRegisters = 1048576

func run(f *os.File) map[string]*Statistics {
	noWorkers := runtime.NumCPU()

	fStat, err := f.Stat()
	if err != nil {
		panic(err)
	}

	var start int64
	var buf [maxBytesPerLine]byte
	ch := make(chan map[string]*Statistics)

	for i := 1; i < noWorkers; i++ {
		finish := fStat.Size() * int64(i) / int64(noWorkers)
		f.ReadAt(buf[:], finish)
		finish += int64(bytes.IndexByte(buf[:], '\n')) + 1

		rIt := io.NewSectionReader(f, start, finish-start)
		go worker(rIt, ch)

		start = finish
	}

	rIt := io.NewSectionReader(f, start, fStat.Size()-start)
	go worker(rIt, ch)

	res := make(map[string]*Statistics)
	for range noWorkers {
		foo := <-ch
		for k, v := range foo {
			if _, ok := res[k]; !ok {
				res[k] = &Statistics{
					Cnt: v.Cnt,
					Min: v.Min,
					Max: v.Max,
					Sum: v.Sum,
				}
			} else {
				res[k].Cnt += v.Cnt
				res[k].Min = min(res[k].Min, v.Min)
				res[k].Max = max(res[k].Max, v.Max)
				res[k].Sum += v.Sum
			}
		}
	}

	return res
}

func worker(
	r io.Reader,
	ch chan<- map[string]*Statistics,
) {
	res := make(map[string]*Statistics, maxCities)
	reader := bufio.NewReader(r)

	for {
		lineIt, err := reader.ReadSlice('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		sepIdx := bytes.IndexByte(lineIt, ';')

		var temperature int
		if lineIt[sepIdx+1] == '-' {
			temperature = -parseDigitsFromBytes(lineIt[sepIdx+2 : len(lineIt)-1])
		} else {
			temperature = parseDigitsFromBytes(lineIt[sepIdx+1 : len(lineIt)-1])
		}

		resIt, ok := res[string(lineIt[:sepIdx])]
		if !ok {
			res[string(lineIt[:sepIdx])] = &Statistics{
				Cnt: 1,
				Max: temperature,
				Min: temperature,
				Sum: temperature,
			}
		} else {
			resIt.Cnt += 1
			resIt.Max = max(resIt.Max, temperature)
			resIt.Min = min(resIt.Min, temperature)
			resIt.Sum += temperature
		}
	}

	ch<-res
}
