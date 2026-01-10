package main

import (
	"fmt"
	"io"
	"slices"

	"golang.org/x/exp/constraints"
)

func write[T constraints.Float](
	w io.Writer,
	stats map[string]*Statistics[T],
) {
	cities := sortedKeys(stats)

	fmt.Fprint(w, "{")
	writeStatistics(
		w,
		cities[0],
		stats[cities[0]],
	)

	for _, cityIt := range cities[1:] {
		fmt.Fprint(w, ", ")
		writeStatistics(
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

func writeStatistics[T constraints.Float](
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
