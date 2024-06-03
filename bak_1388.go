package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bak1388() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	graph := make([][]rune, n)
	for i := 0; i < n; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		graph[i] = []rune(str)
	}

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	var dfs func(int, int, rune, int, int)
	dfs = func(x, y int, char rune, dx, dy int) {
		visited[x][y] = true
		nx, ny := x+dx, y+dy
		if nx >= 0 && nx < n && ny >= 0 && ny < m && !visited[nx][ny] && graph[nx][ny] == char {
			dfs(nx, ny, char, dx, dy)
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] {
				count++
				if graph[i][j] == '-' {
					dfs(i, j, '-', 0, 1) // horizontal direction
				} else if graph[i][j] == '|' {
					dfs(i, j, '|', 1, 0) // vertical direction
				}
			}
		}
	}

	fmt.Fprintln(writer, count)
}

func main() {
	bak1388()
}
