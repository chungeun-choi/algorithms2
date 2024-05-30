package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func bak2230() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	dataList := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &dataList[i])
	}

	// 데이터 정렬
	sort.Ints(dataList)

	leastValue := math.MaxInt64
	i, j := 0, 1

	for i < n && j < n {
		value := dataList[j] - dataList[i]
		if value < m {
			j++
		} else {
			if leastValue > value {
				leastValue = value
			}
			i++
			if i == j {
				j++
			}
		}
	}

	if leastValue == math.MaxInt64 {
		leastValue = 0
	}

	fmt.Fprintln(writer, leastValue)
}
