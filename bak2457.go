package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Flower struct {
	start, end int
}

func bak2457() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	flowers := make([]Flower, n)
	for i := 0; i < n; i++ {
		var m1, d1, m2, d2 int
		fmt.Fscanf(reader, "%d %d %d %d\n", &m1, &d1, &m2, &d2)
		start := m1*100 + d1
		end := m2*100 + d2
		flowers[i] = Flower{start, end}
	}

	// Sort flowers by start date, and by end date if start dates are the same
	sort.Slice(flowers, func(i, j int) bool {
		if flowers[i].start == flowers[j].start {
			return flowers[i].end < flowers[j].end
		}
		return flowers[i].start < flowers[j].start
	})

	const startDate = 301
	const endDate = 1130

	count := 0
	currentEnd := startDate
	nextEnd := startDate
	idx := 0

	for {
		hasFlower := false
		for idx < n && flowers[idx].start <= currentEnd {
			if flowers[idx].end > nextEnd {
				nextEnd = flowers[idx].end
				hasFlower = true
			}
			idx++
		}
		if !hasFlower {
			break
		}
		currentEnd = nextEnd
		count++
		if currentEnd > endDate {
			fmt.Fprintf(writer, "%d\n", count)
			return
		}
	}

	fmt.Fprintf(writer, "0\n")
}
