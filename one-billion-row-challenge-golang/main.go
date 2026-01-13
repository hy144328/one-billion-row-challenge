package main

import (
	"bufio"
	"bytes"
	"io"
	"math"
	"strconv"
	"strings"
)

const maxCities = 10000
const maxLine = 128
const noRegisters = 1048576

func run(r io.Reader) map[string]*Statistics[float64] {
	scanner := bufio.NewScanner(r)

	counts := make(map[string]int, maxCities)
	maxs := make(map[string]float64, maxCities)
	mins := make(map[string]float64, maxCities)
	sums := make(map[string]float64, maxCities)

	for scanner.Scan() {
		lineIt := scanner.Text()
		words := strings.Split(lineIt, ";")

		city := words[0]
		temperature, err := strconv.ParseFloat(words[1], 64)
		if err != nil {
			panic(err)
		}

		counts[city] += 1
		sums[city] += temperature

		if counts[city] == 1 {
			maxs[city] = temperature
			mins[city] = temperature
		} else {
			maxs[city] = max(maxs[city], temperature)
			mins[city] = min(mins[city], temperature)
		}
	}

	res := make(map[string]*Statistics[float64], len(counts))

	for cityIt, countIt := range counts {
		res[cityIt] = &Statistics[float64]{
			Cnt: countIt,
			Max: maxs[cityIt],
			Min: mins[cityIt],
			Sum: sums[cityIt],
		}
	}

	return res
}

func run1(r io.Reader) map[string]*Statistics[float64] {
	res := make(map[string]*Statistics[float64], maxCities)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lineIt := scanner.Text()
		words := strings.Split(lineIt, ";")

		city := words[0]
		temperature, err := strconv.ParseFloat(words[1], 64)
		if err != nil {
			panic(err)
		}

		resIt, ok := res[city]
		if !ok {
			resIt = &Statistics[float64]{
				Cnt: 1,
				Max: temperature,
				Min: temperature,
				Sum: temperature,
			}
			res[city] = resIt
		} else {
			resIt.Cnt += 1
			resIt.Max = max(resIt.Max, temperature)
			resIt.Min = min(resIt.Min, temperature)
			resIt.Sum += temperature
		}
	}

	return res
}

func run2(r io.Reader) map[string]*Statistics[int] {
	res := make(map[string]*Statistics[int], maxCities)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lineIt := scanner.Text()
		words := strings.Split(lineIt, ";")

		city := words[0]
		word1 := words[1]
		word1len := len(word1)

		temperature10, err := strconv.Atoi(word1[:word1len-2])
		if err != nil {
			panic(err)
		}

		sgn := 1
		if word1[0] == '-' {
			sgn = -1
		}

		temperature1 := word1[word1len-1] - '0'
		temperature := 10 * temperature10 + sgn * int(temperature1)

		resIt, ok := res[city]
		if !ok {
			resIt = &Statistics[int]{
				Cnt: 1,
				Max: temperature,
				Min: temperature,
				Sum: temperature,
			}
			res[city] = resIt
		} else {
			resIt.Cnt += 1
			resIt.Max = max(resIt.Max, temperature)
			resIt.Min = min(resIt.Min, temperature)
			resIt.Sum += temperature
		}
	}

	return res
}

func run3(r io.Reader) map[string]*Statistics[int] {
	res := make(map[string]*Statistics[int], maxCities)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lineIt := scanner.Text()
		lineItLen := len(lineIt)
		sepIdx := strings.IndexByte(lineIt, ';')

		city := lineIt[:sepIdx]

		temperature10, err := strconv.Atoi(lineIt[sepIdx+1:lineItLen-2])
		if err != nil {
			panic(err)
		}

		sgn := 1
		if lineIt[sepIdx+1] == '-' {
			sgn = -1
		}

		temperature1 := lineIt[lineItLen-1] - '0'
		temperature := 10 * temperature10 + sgn * int(temperature1)

		resIt, ok := res[city]
		if !ok {
			resIt = &Statistics[int]{
				Cnt: 1,
				Max: temperature,
				Min: temperature,
				Sum: temperature,
			}
			res[city] = resIt
		} else {
			resIt.Cnt += 1
			resIt.Max = max(resIt.Max, temperature)
			resIt.Min = min(resIt.Min, temperature)
			resIt.Sum += temperature
		}
	}

	return res
}

func run4(r io.Reader) map[string]*Statistics[int] {
	res := make(map[string]*Statistics[int], maxCities)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lineIt := scanner.Bytes()
		lineItLen := len(lineIt)
		sepIdx := bytes.IndexByte(lineIt, ';')

		city := lineIt[:sepIdx]

		temperature10, err := strconv.Atoi(string(lineIt[sepIdx+1:lineItLen-2]))
		if err != nil {
			panic(err)
		}

		sgn := 1
		if lineIt[sepIdx+1] == '-' {
			sgn = -1
		}

		temperature1 := lineIt[lineItLen-1] - '0'
		temperature := 10 * temperature10 + sgn * int(temperature1)

		resIt, ok := res[string(city)]
		if !ok {
			resIt = &Statistics[int]{
				Cnt: 1,
				Max: temperature,
				Min: temperature,
				Sum: temperature,
			}
			res[string(city)] = resIt
		} else {
			resIt.Cnt += 1
			resIt.Max = max(resIt.Max, temperature)
			resIt.Min = min(resIt.Min, temperature)
			resIt.Sum += temperature
		}
	}

	return res
}

func run5(r io.Reader) map[string]*Statistics[int] {
	res := make(map[string]*Statistics[int], maxCities)
	reader := bufio.NewReader(r)

	for {
		city, err := reader.ReadSlice(';')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		city = city[:len(city)-1]
		resIt, ok := res[string(city)]
		if !ok {
			resIt = &Statistics[int]{
				Max: math.MinInt,
				Min: math.MaxInt,
			}
			res[string(city)] = resIt
		}

		temperature := 0
		sgn := 1

		for {
			b, err := reader.ReadByte()
			if err != nil {
				panic(err)
			}

			if b >= '0' && b <= '9' {
				temperature *= 10
				temperature += int(b - '0')
			} else if b == '\n' {
				break
			} else if b == '-' {
				sgn = -1
			}
		}

		temperature *= sgn
		resIt.Cnt += 1
		resIt.Max = max(resIt.Max, temperature)
		resIt.Min = min(resIt.Min, temperature)
		resIt.Sum += temperature
	}

	return res
}

func run6(r io.Reader) map[string]*Statistics[int] {
	var city []byte
	var temperature int

	res := make(map[string]*Statistics[int], maxCities)
	scanner := bufio.NewScanner(r)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if !atEOF && len(data) < maxLine {
			return 0, nil, nil
		}

		sepIdx := bytes.IndexByte(data, ';')
		city = data[0:sepIdx]
		temp := 0
		sgn := 1

		for i := sepIdx + 1; i < len(data); i++ {
			b := data[i]

			if b >= '0' && b <= '9' {
				temp *= 10
				temp += int(b - '0')
			} else if b == '\n' {
				temperature = sgn * temp
				return i + 1, data[sepIdx+1:i], nil
			} else if b == '-' {
				sgn = -1
			}
		}

		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	})

	for scanner.Scan() {
		resIt, ok := res[string(city)]
		if !ok {
			resIt = &Statistics[int]{
				Max: math.MinInt,
				Min: math.MaxInt,
			}
			res[string(city)] = resIt
		}

		resIt.Cnt += 1
		resIt.Max = max(resIt.Max, temperature)
		resIt.Min = min(resIt.Min, temperature)
		resIt.Sum += temperature
	}

	return res
}

func run7(r io.Reader) *BytesMap[Statistics[int]] {
	var city []byte
	var temperature int
	var h uint32

	res := NewBytesMap[Statistics[int]](noRegisters)
	scanner := bufio.NewScanner(r)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if !atEOF && len(data) < maxLine {
			return 0, nil, nil
		}

		var sepIdx int
		var fnv uint32 = 2166136261

		for i, b := range data {
			if b == ';' {
				sepIdx = i
				h = fnv
				break
			}

			fnv ^= uint32(b)
			fnv *= 16777619
		}

		city = data[0:sepIdx]
		temp := 0
		sgn := 1

		for i := sepIdx + 1; i < len(data); i++ {
			b := data[i]

			if b >= '0' && b <= '9' {
				temp *= 10
				temp += int(b - '0')
			} else if b == '\n' {
				temperature = sgn * temp
				return i + 1, data[sepIdx+1:i], nil
			} else if b == '-' {
				sgn = -1
			}
		}

		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, nil
	})

	for scanner.Scan() {
		resIt, ok := res.GetOrCreate(city, h)

		if !ok {
			resIt.Cnt = 1
			resIt.Max = temperature
			resIt.Min = temperature
			resIt.Sum = temperature
		} else {
			resIt.Cnt += 1
			resIt.Max = max(resIt.Max, temperature)
			resIt.Min = min(resIt.Min, temperature)
			resIt.Sum += temperature
		}
	}

	return res
}
