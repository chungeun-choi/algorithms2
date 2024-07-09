package baekjoon

import (
	"bufio"
	"fmt"
	"os"
)

func dfs5567(start int, graph map[int][]int, visited []bool, depth int) int {
	var result int

	visited[start] = true
	depth += 1
	if depth >= 3 {
		return 1
	}

	for _, value := range graph[start] {
		if !visited[value] {
			result += dfs5567(value, graph, visited, depth)
		}
	}

	return result
}

func Baek5567() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var (
		n, listLenth int
		graph        map[int][]int = make(map[int][]int)
	)

	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &listLenth)

	for i := 0; i < listLenth; i++ {
		var root, leaf int
		fmt.Fscanf(reader, "%d %d\n", &root, &leaf)

		if _, ok := graph[root]; ok {
			graph[root] = append(graph[root], leaf)
		} else {
			graph[root] = []int{leaf}
		}

	}

	if _, ok := graph[1]; !ok {
		fmt.Fprintln(writer, 0)
		return
	}

	fmt.Fprintln(writer, dfs5567(1, graph, make([]bool, n), 0))

}
