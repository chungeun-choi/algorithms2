package baekjoon

import "fmt"

/*
1. 범위내의 최대 높이를 확인 -> 강수량을 최대 높이까지만 연산하기위해서
2. 탐색 알고리즘임에 따라 visited 2차원 슬라이스를 생성
3. dfs 알고리즘을 통해서 시작지점 (0,0) 부터 탐색 시작 -> dfs 한번 돌 때마다 safeZone 갯수 증가
 3-1. dfs 알고리즘에서 이동 반경은 상하좌우
 3-2. dfs 알고리즘에서 높이가 강수량보다 높은 부분만 visited 에 체크
4. safeZone 값이 maxSafeZone 보다 높을 경우 갱신
*/

func Back2468() int {
	var N int
	var graph [][]int
	var visited [][]bool
	var maxHeight int = 0
	var maxSafeZones int = 1
	var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	fmt.Scan(&N)
	graph = make([][]int, N)

	for i := 0; i < N; i++ {
		graph[i] = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&graph[i][j])
			if graph[i][j] > maxHeight {
				maxHeight = graph[i][j]
			}
		}
	}

	for h := 1; h < maxHeight; h++ {
		visited = make([][]bool, N)
		for i := 0; i < N; i++ {
			visited[i] = make([]bool, N)
		}

		safeZones := 0
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if graph[i][j] > h && !visited[i][j] {
					func(x, y, h int) {
						stack := make([][2]int, 0)
						visited[i][j] = true
						stack = append(stack, [2]int{x, y})

						for len(stack) > 0 {
							curr := stack[len(stack)-1]
							currX := curr[0]
							currY := curr[1]
							stack = stack[:len(stack)-1]

							for _, dir := range directions {
								dx := currX - dir[0]
								dy := currY - dir[1]

								if 0 <= dx && dx < N && 0 <= dy && dy < N && !visited[dx][dy] && graph[dx][dy] > h {
									visited[dx][dy] = true
									stack = append(stack, [2]int{dx, dy})
								}
							}

						}

					}(i, j, h)
					safeZones++
				}
			}
		}

		if safeZones > maxSafeZones {
			maxSafeZones = safeZones
		}

	}

	return maxSafeZones
}
