package main

import (
	"fmt"
	"io"
	"slices"
)

func writeInt(
	w io.Writer,
	stats map[string]*Statistics,
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

func writeIntStatistics(
	w io.Writer,
	city string,
	stats *Statistics,
) {
	fmt.Fprint(w, city)
	fmt.Fprint(w, "=")
	fmt.Fprintf(w, "%.1f", 0.1*float32(stats.Min))
	fmt.Fprint(w, "/")
	fmt.Fprintf(w, "%.1f", 0.1*float64(stats.Sum)/float64(stats.Cnt))
	fmt.Fprint(w, "/")
	fmt.Fprintf(w, "%.1f", 0.1*float32(stats.Max))
}
