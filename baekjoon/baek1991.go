package baekjoon

import (
	"bufio"
	"fmt"
	"os"
)

func preorder(node string, graph map[string][2]string, writer *bufio.Writer) {
	if node == "." {
		return
	}
	fmt.Fprint(writer, node)
	preorder(graph[node][0], graph, writer)
	preorder(graph[node][1], graph, writer)
}

func inorder(node string, graph map[string][2]string, writer *bufio.Writer) {
	if node == "." {
		return
	}
	inorder(graph[node][0], graph, writer)
	fmt.Fprint(writer, node)
	inorder(graph[node][1], graph, writer)
}

func postorder(node string, graph map[string][2]string, writer *bufio.Writer) {
	if node == "." {
		return
	}
	postorder(graph[node][0], graph, writer)
	postorder(graph[node][1], graph, writer)
	fmt.Fprint(writer, node)
}

func Baek1991() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	graph := make(map[string][2]string, n)

	for i := 0; i < n; i++ {
		var root, left, right string
		fmt.Fscanf(reader, "%s %s %s\n", &root, &left, &right)
		graph[root] = [2]string{left, right}
	}

	preorder("A", graph, writer)
	fmt.Fprintln(writer)
	inorder("A", graph, writer)
	fmt.Fprintln(writer)
	postorder("A", graph, writer)
	fmt.Fprintln(writer)
}
