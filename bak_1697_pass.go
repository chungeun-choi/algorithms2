package main

var (
	visited   [][]bool
	pictures  int
	maxSize   int
	positions [][2]int = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	N, K      int
	graph     [][]int
)

func bfstest(x, y int) int {
	var (
		result int = 1
		queue  [][2]int
	)

	queue = append(queue, [2]int{x, y})
	visited[x][y] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		currX, currY := curr[0], curr[1]

		for _, pos := range positions {
			dx, dy := currX+pos[0], currY+pos[1]

			if 0 <= dx && dx < N && 0 <= dy && dy < K && !visited[dx][dy] && graph[dx][dy] == 1 {
				visited[dx][dy] = true
				queue = append(queue, [2]int{dx, dy})
				result++
			}
		}
	}
	return result
}

func back1926(N, M int, graph [][]int) [2]int {
	visited = make([][]bool, N)
	for i := 0; i < N; i++ {
		visited[i] = make([]bool, M)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if !visited[i][j] && graph[i][j] == 1 {
				result := bfstest(i, j)
				pictures++
				if maxSize < result {
					maxSize = result
				}
			}
		}
	}
	return [2]int{pictures, maxSize}
}
