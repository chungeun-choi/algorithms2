package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func bak1183() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	differences := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := scanner.Text()
		ab := make([]int, 2)
		for j, s := range input {
			ab[j], _ = strconv.Atoi(string(s))
		}
		a, b := ab[0], ab[1]
		differences[i] = a - b
	}

	sort.Ints(differences)

	var median int
	if n%2 == 0 {
		median = (differences[n/2-1] + differences[n/2]) / 2
	} else {
		median = differences[n/2]
	}

	fmt.Fprintln(writer, median)
}
