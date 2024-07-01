package main

import (
	"algorthms/baekjoon"
	"algorthms/base/dijkstra"
	"algorthms/programers"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func callBak1697() {
	var N, K int
	fmt.Scan(&N, &K)
	fmt.Println(baekjoon.Bfs(N, K))
}

func callBack2468() {
	fmt.Println(baekjoon.Back2468())
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

	result := baekjoon.Back1926(N, K, graph)

	fmt.Printf("%v \n%v\n", result[0], result[1])
}

func callBack10773() {
	fmt.Println(baekjoon.Bak10773())
}

func callBack1874() {
	baekjoon.Bak1874()
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

func solved_programers_parking() {
	var (
		fees []int = []int{
			180, 5000, 10, 600,
		}
		records []string = []string{
			"05:34 5961 IN", "06:00 0000 IN", "06:34 0000 OUT", "07:59 5961 OUT", "07:59 0148 IN", "18:59 0000 IN", "19:09 0148 OUT", "22:59 5961 IN", "23:00 5961 OUT",
		}
	)
	fmt.Println(programers.Solution3(fees, records))

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
	//baek1182()
	//baek15651()
	//baek15652()
	//baek6198()
	//baek3015()
	//baek22866()
	//baek10807()
	//baek1629()
	//baek11729()
	//baekjoon.Baek1991()
	//fmt.Println(programers.Solution1("027778888"))
	//fmt.Println(
	//	programers.Solution2([][]string{
	//		{
	//			"yellow_hat", "headgear",
	//		},
	//		{
	//			"blue_sunglasses", "eyewear",
	//		},
	//		{
	//			"green_turban", "headgear",
	//		},
	//	}),
	//)
	//solved_programers_parking()
	//baekjoon.Baek14425()
	//baekjoon.Baek5052()
	//baekjoon.Baek7432()
	//baekjoon.Baek10816()
	//baekjoon.Baek1260()
	//baekjoon.Baek2164()
	baekjoon.Baek7662()
}
