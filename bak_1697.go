package main

func bfs(n, k int) int {
	const max = 100001
	queue := make([][2]int, max)
	visited := make([]bool, max)

	visited[n] = true
	queue = append(queue, [2]int{n, 0})

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		currLoc := curr[0]
		currTime := curr[1]

		if currLoc == k {
			return currTime
		}

		positions := []int{currLoc - 1, currLoc + 1, currLoc * 2}

		for _, pos := range positions {
			if 0 <= pos && !visited[pos] {
				visited[pos] = true
				queue = append(queue, [2]int{pos, currTime + 1})
			}
		}

	}
	return -1
}
