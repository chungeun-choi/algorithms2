package dijkstra

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Edge struct {
	node int
	cost int
}

func MakeGraph() ([][]int, int) {
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

func Dijkstra(graph [][]int, start int) ([]int, []int) {
	n := len(graph)
	dist := make([]int, n)
	prev := make([]int, n)
	for i := range dist {
		dist[i] = 1e9 // 무한대로 초기화
		prev[i] = -1  // 이전 노드를 -1로 초기화
	}
	dist[start] = 0

	queue := []Edge{{node: start, cost: 0}}

	for len(queue) > 0 {
		// sort the queue based on cost
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].cost < queue[j].cost
		})

		// pop the first element
		u := queue[0]
		queue = queue[1:]
		uNode := u.node
		uCost := u.cost

		if uCost > dist[uNode] {
			continue
		}

		for v := 0; v < n; v++ {
			if graph[uNode][v] != 0 {
				alt := uCost + graph[uNode][v]
				if alt < dist[v] {
					dist[v] = alt
					prev[v] = uNode // 이전 노드 기록
					queue = append(queue, Edge{node: v, cost: alt})
				}
			}
		}
	}
	return dist, prev
}

func PrintPath(prev []int, target int) {
	if prev[target] != -1 {
		PrintPath(prev, prev[target])
	}
	fmt.Print(target, " ")
}
