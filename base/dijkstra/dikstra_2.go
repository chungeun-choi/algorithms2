package dijkstra

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func MakeGraph2() (int, [][]int) {
	reader := bufio.NewReader(os.Stdin)
	var (
		n     int
		graph [][]int
	)

	fmt.Fscanln(reader, &n)

	graph = make([][]int, n)

	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		strSlice := strings.Split(str, " ")
		for j, s := range strSlice {
			graph[i][j], _ = strconv.Atoi(s)
		}
	}

	return n, graph
}

type Edge2 struct {
	Start    int
	Priority int
}

func Dijkstra2(start int) ([]int, []int) {
	n, graph := MakeGraph2()

	distance := make([]int, n)
	prev := make([]int, n)
	queue := make([]Edge2, 0, n)

	for i := 0; i < n; i++ {
		distance[i] = 1e9
		prev[i] = -1
	}

	queue = append(queue, Edge2{start, 0})
	prev[start] = 0

	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].Priority < queue[j].Priority
		})

		curr := queue[0]
		queue = queue[1:]
		node := curr.Start
		priority := curr.Priority

		if distance[node] < priority {
			continue
		}

		for i := 0; i < n; i++ {
			if graph[node][i] != 0 {
				sumPriority := priority + graph[node][i]
				if sumPriority < distance[node] {
					distance[i] = sumPriority
					prev[i] = node
					queue = append(queue, Edge2{i, sumPriority})
				}
			}
		}
	}
	return distance, prev
}
