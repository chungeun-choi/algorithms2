package baekjoon

import (
	"bufio"
	"fmt"
	"os"
)

func bfs(start int, graph map[int][]int, visited []bool, writer *bufio.Writer) {
	var queue []int
	queue = append(queue, start)
	visited[start-1] = true

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]
		fmt.Fprint(writer, currNode) // 줄바꿈 추가

		for _, value := range graph[currNode] {
			if !visited[value-1] {
				queue = append(queue, value)
				visited[value-1] = true // 방문 여부 업데이트
			}
		}
	}
}

func dfs(start int, graph map[int][]int, visited []bool, writer *bufio.Writer) {
	visited[start-1] = true
	fmt.Fprint(writer, start)

	for _, value := range graph[start] {
		if !visited[value-1] {
			dfs(value, graph, visited, writer)
		}
	}

}

func Baek1260() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var (
		n, m, v int
	)

	fmt.Fscanf(reader, "%d %d %d\n", &n, &m, &v)

	graph := make(map[int][]int)

	for i := 0; i < m; i++ {
		var root, leaf int
		fmt.Fscanf(reader, "%d %d\n", &root, &leaf)

		if _, ok := graph[root]; !ok {
			graph[root] = []int{leaf}
		} else {
			graph[root] = append(graph[root], leaf)
		}
	}

	visited := make([]bool, m)
	dfs(v, graph, visited, writer)

	fmt.Fprintln(writer)

	visited = make([]bool, m)
	bfs(v, graph, visited, writer)

}
