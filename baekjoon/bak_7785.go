package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func bak7785() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var enterList [][2]string

	for i := 0; i < n; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		strs := strings.Split(text, " ")

		enterList = append(enterList, [2]string{strs[0], strs[1]})
	}

	var enterMap map[string]int = make(map[string]int)

	for _, enterObj := range enterList {
		if _, ok := enterMap[enterObj[0]]; ok && enterObj[1] == "leave" {
			delete(enterMap, enterObj[0])
		} else {
			enterMap[enterObj[0]] = 1
		}
	}

	var remaining []string
	for key := range enterMap {
		remaining = append(remaining, key)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(remaining)))

	for _, name := range remaining {
		fmt.Fprintln(writer, name)
	}
}
