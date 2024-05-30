package main

func back1926Wrong(N, M int, graph [][]int) [2]int {
	var (
		visited  [][]bool
		pictures int
		maxSize  int
	)

	dfs := func(x, y int) int {
		var (
			result    int      = 1
			positions [][2]int = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
			stack     [][2]int
		)

		stack = append(stack, [2]int{x, y})
		visited[x][y] = true

		for len(stack) > 0 {

			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			currX, currY := curr[0], curr[1]

			for _, pos := range positions {
				dx, dy := currX+pos[0], currY+pos[1]

				if 0 <= dx && dx < N && 0 <= dy && dy < M && !visited[dx][dy] && graph[dx][dy] == 1 {
					visited[dx][dy] = true
					stack = append(stack, [2]int{dx, dy})
					result++
				}
			}
		}
		return result
	}

	visited = make([][]bool, N)
	for i := 0; i < N; i++ {
		visited[i] = make([]bool, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if !visited[i][j] && graph[i][j] == 1 {
				result := dfs(i, j)
				pictures++
				if maxSize < result {
					maxSize = result
				}
			}
		}
	}

	return [2]int{pictures, maxSize}
}
