package main

import (
	"fmt"
	"io"
	"slices"

	"golang.org/x/exp/constraints"
)

func writeFloat[T constraints.Float](
	w io.Writer,
	stats map[string]*Statistics[T],
) {
	cities := sortedKeys(stats)

	fmt.Fprint(w, "{")
	if len(stats) == 0 {
		fmt.Fprint(w, "}\n")
		return
	}

	writeFloatStatistics(
		w,
		cities[0],
		stats[cities[0]],
	)

	for _, cityIt := range cities[1:] {
		fmt.Fprint(w, ", ")
		writeFloatStatistics(
			w,
			cityIt,
			stats[cityIt],
		)
	}

	fmt.Fprint(w, "}\n")
}

func writeInt[T constraints.Integer](
	w io.Writer,
	stats map[string]*Statistics[T],
) {
	cities := sortedKeys(stats)

	fmt.Fprint(w, "{")
	if len(stats) == 0 {
		fmt.Fprint(w, "}\n")
		return
	}

	writeIntStatistics(
		w,
		cities[0],
		stats[cities[0]],
	)

	for _, cityIt := range cities[1:] {
		fmt.Fprint(w, ", ")
		writeIntStatistics(
			w,
			cityIt,
			stats[cityIt],
		)
	}

	fmt.Fprint(w, "}\n")
}

func sortedKeys[T any](m map[string]T) []string {
	res := make([]string, 0, len(m))

	for k := range m {
		res = append(res, k)
	}

	slices.Sort(res)
	return res
}

func writeFloatStatistics[T constraints.Float](
	w io.Writer,
	city string,
	stats *Statistics[T],
) {
	fmt.Fprint(w, city)
	fmt.Fprint(w, "=")
	fmt.Fprintf(w, "%.1f", stats.Min)
	fmt.Fprint(w, "/")
	fmt.Fprintf(w, "%.1f", stats.Sum / T(stats.Cnt))
	fmt.Fprint(w, "/")
	fmt.Fprintf(w, "%.1f", stats.Max)
}

func writeIntStatistics[T constraints.Integer](
	w io.Writer,
	city string,
	stats *Statistics[T],
) {
	fmt.Fprint(w, city)
	fmt.Fprint(w, "=")
	fmt.Fprintf(w, "%.1f", 0.1 * float32(stats.Min))
	fmt.Fprint(w, "/")
	fmt.Fprintf(w, "%.1f", 0.1 * float64(stats.Sum) / float64(stats.Cnt))
	fmt.Fprint(w, "/")
	fmt.Fprintf(w, "%.1f", 0.1 * float32(stats.Max))
}
