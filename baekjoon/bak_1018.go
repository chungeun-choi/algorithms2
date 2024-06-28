package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bak1018() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	var graph [][]rune

	graph = make([][]rune, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]rune, m)
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)

		for j, value := range str {
			graph[i][j] = value
		}
	}

	minChangeCount := 1<<31 - 1

	for i := 0; i <= n-8; i++ {
		for j := 0; j <= m-8; j++ {
			countW := countChanges(graph, i, j, 'W')
			countB := countChanges(graph, i, j, 'B')
			if minChangeCount > countW {
				minChangeCount = countW
			}
			if minChangeCount > countB {
				minChangeCount = countB
			}
		}
	}

	fmt.Fprintln(writer, minChangeCount)
}

func countChanges(graph [][]rune, x int, y int, startColor rune) int {
	var changeCount int
	colors := [2]rune{'W', 'B'}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			expectedColor := colors[(i+j)%2]
			if startColor == 'B' {
				expectedColor = colors[(i+j+1)%2]
			}
			if graph[x+i][y+j] != expectedColor {
				changeCount++
			}
		}
	}
	return changeCount
}
