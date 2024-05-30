package dijkstra

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EdgeMe struct {
	Start    int
	Priority int
}

func MakeGraphMe() ([][]int, int) {
	var (
		graph [][]int
		n     int
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of nodes: ")
	fmt.Fscanln(reader, &n)
	graph = make([][]int, n)

	fmt.Println("Enter the adjacency matrix:")
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		strSlice := strings.Split(str, " ")
		for j, s := range strSlice {
			graph[i][j], _ = strconv.Atoi(s)
		}
	}
	return graph, n
}

func DijkstraMe(start int) {
	graph, n := MakeGraphMe()
	// 탐색 알고리즘에서 꼭 필요한 내용은 방문한 노드를 담을 자료구조와 visited 테이블
	queue := make([]EdgeMe, 0, n)
	visited := make([]int, n)
	prev := make([]int, n)
	for i := 0; i < n; i++ {
		visited[i] = 1e9
		prev[i] = -1
	}
	// queue 초기 값 초기화
	queue = append(queue, EdgeMe{start, 0})
	visited[start] = 0

	for len(queue) > 0 {
		// queue 에서 데이터 pop 하기 전 최소 값 기준으로 정렬 (priority queue 구현 없이 정렬로 진행)
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].Priority < queue[j].Priority
		})

		curr := queue[0]
		queue = queue[1:]
		priority := curr.Priority
		node := curr.Start

		if visited[node] < priority {
			continue
		}

		for i := 0; i < n; i++ {
			if graph[node][i] != 0 {
				sumPriority := graph[node][i] + priority
				// 현재까지의 거리를 계산하여 visited 테이블에 저장되어진 최소 거리와 비교하여 확인
				if sumPriority < visited[node] {
					queue = append(queue, EdgeMe{node, sumPriority})
					// 방분한 노드에 대해서 값을 변경 시켜줘야됨
					visited[i] = sumPriority
					prev[i] = node
				}
			}
		}

	}

}
