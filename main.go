package main

import (
	"algorthms/base/dijkstra"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func callBak1697() {
	var N, K int
	fmt.Scan(&N, &K)
	fmt.Println(bfs(N, K))
}

func callBack2468() {
	fmt.Println(back2468())
}

func callBack1926() {
	var N, K int
	var graph [][]int

	fmt.Scan(&N, &K)
	graph = make([][]int, N)

	for i := 0; i < N; i++ {
		graph[i] = make([]int, K)
		for j := 0; j < K; j++ {
			fmt.Scan(&graph[i][j])
		}
	}

	result := back1926(N, K, graph)

	fmt.Printf("%v \n%v\n", result[0], result[1])
}

func callBack10773() {
	fmt.Println(bak10773())
}

func callBack1874() {
	bak1874()
}

func callBaseDijkstra() {
	graph, _ := dijkstra.MakeGraph()
	fmt.Print("Enter the start node: ")
	reader := bufio.NewReader(os.Stdin)
	startStr, _ := reader.ReadString('\n')
	startStr = strings.TrimSpace(startStr)
	start, _ := strconv.Atoi(startStr)

	dist, prev := dijkstra.Dijkstra(graph, start)
	fmt.Println("Shortest distances from start node:", dist)
	fmt.Println("Paths from start node:")

	for i := range dist {
		if i != start {
			fmt.Printf("Path to %d: ", i)
			dijkstra.PrintPath(prev, i)
			fmt.Printf("(cost: %d)\n", dist[i])
		}
	}
}

func main() {
	//callBak1697()
	//callBack2468()
	//callBack1926()
	//callBack10773()
	//callBack1874()
	//back2493()
	//bak11728()
	//bak2577()
	//bak2457()
	//callBaseDijkstra()
	//bak1920()
	//bak2230()
	//bak1806()
	//bak7785()
	//bak1015()
	//bak1018()
	//bak1026()
	//bak1388()
	//bak1063()
	//bak1515()
	//bak2193()
	//bak1966()
	//bak2295()
	//baek1092()
	//baek1245()
	//baek1469()
	//baek1501()
	//baek9663()
	baek1182()
}
