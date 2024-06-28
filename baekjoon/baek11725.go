package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func baek11725() {
	var n int
	fmt.Scan(&n)

	tree := make([][]int, n+1)
	parents := make([]int, n+1)
	visited := make([]bool, n+1)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < n-1; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		nodes := strings.Split(line, " ")
		a, _ := strconv.Atoi(nodes[0])
		b, _ := strconv.Atoi(nodes[1])
		tree[a] = append(tree[a], b)
		tree[b] = append(tree[b], a)
	}

	dfs11725(1, tree, parents, visited)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for i := 2; i <= n; i++ {
		fmt.Fprintln(writer, parents[i])
	}
}

func dfs11725(node int, tree [][]int, parents []int, visited []bool) {
	visited[node] = true
	for _, neighbor := range tree[node] {
		if !visited[neighbor] {
			parents[neighbor] = node
			dfs11725(neighbor, tree, parents, visited)
		}
	}
}
