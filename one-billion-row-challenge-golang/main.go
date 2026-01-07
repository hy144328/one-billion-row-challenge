package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("measurements_6.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	counts := make(map[string]int, 10000)
	maxs := make(map[string]float64, 10000)
	mins := make(map[string]float64, 10000)
	sums := make(map[string]float64, 10000)

	for scanner.Scan() {
		line_it := scanner.Text()
		words := strings.Split(line_it, ";")

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

	cities := sortedKeys(counts)
	res0 := make([]string, len(cities))

	for cityCt, cityIt := range cities {
		res0[cityCt] = fmt.Sprintf(
			"%s=%.1f/%.1f/%.1f",
			cityIt,
			mins[cityIt],
			sums[cityIt] / float64(counts[cityIt]),
			maxs[cityIt],
		)
	}

	fmt.Println("{" + strings.Join(res0, ", ") + "}")
}

func sortedKeys[T any](m map[string]T) []string {
	res := make([]string, 0, len(m))

	for k := range m {
		res = append(res, k)
	}

	slices.Sort(res)
	return res
}
