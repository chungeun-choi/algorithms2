package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type DIR struct {
	data  string
	child []*DIR
}

func NewDIR(data string) *DIR {
	return &DIR{data: data, child: []*DIR{}}
}

func (d *DIR) update(path string) {
	if path == "" {
		return
	}
	splitIndex := strings.Index(path, "\\")
	var input string
	if splitIndex == -1 {
		input = path
		path = ""
	} else {
		input = path[:splitIndex]
		path = path[splitIndex+1:]
	}

	var child *DIR
	for _, c := range d.child {
		if c.data == input {
			child = c
			break
		}
	}
	if child == nil {
		child = NewDIR(input)
		d.child = append(d.child, child)
	}

	child.update(path)
}

func (d *DIR) dfs(depth int, writer *bufio.Writer) {
	if depth != -1 {
		writer.WriteString(strings.Repeat(" ", depth))
		writer.WriteString(d.data + "\n")
	}
	sort.Slice(d.child, func(i, j int) bool {
		return d.child[i].data < d.child[j].data
	})
	for _, c := range d.child {
		c.dfs(depth+1, writer)
	}
}

func Baek7432() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)
	root := NewDIR("")

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		root.update(line)
	}

	root.dfs(-1, writer)
}
