package main

import (
	"bufio"
	"bytes"
	"io"
)

const maxCities = 10000

func run(r io.Reader) map[string]*Statistics {
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
				Max: int16(temperature),
				Min: int16(temperature),
				Sum: int64(temperature),
			}
		} else {
			resIt.Cnt += 1
			resIt.Max = max(resIt.Max, int16(temperature))
			resIt.Min = min(resIt.Min, int16(temperature))
			resIt.Sum += int64(temperature)
		}
	}

	return res
}
