package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bak1806() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, s int
	fmt.Fscanln(reader, &n, &s)

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	strs := strings.Split(text, " ")

	nList := make([]int, n)
	for i, str := range strs {
		nList[i], _ = strconv.Atoi(str)
	}

	result := n + 1
	sumValue := 0
	start := 0

	for end := 0; end < n; end++ {
		sumValue += nList[end]

		for sumValue >= s {
			if end-start+1 < result {
				result = end - start + 1
			}
			sumValue -= nList[start]
			start++
		}
	}

	if result == n+1 {
		result = 0
	}

	fmt.Fprintln(writer, result)
}
