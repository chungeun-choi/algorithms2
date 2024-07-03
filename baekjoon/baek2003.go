package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Baek2003() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int

	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	numbers := make([]int, n)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Fields(input)

	for i, value := range words {
		number, _ := strconv.Atoi(value)
		numbers[i] = number
	}

	// 슬라이딩 윈도우를 사용하여 연속 부분 수열의 합이 m인 경우의 수 찾기
	count := 0
	sum := 0
	start := 0

	for end := 0; end < n; end++ {
		sum += numbers[end]

		for sum > m {
			sum -= numbers[start]
			start++
		}

		if sum == m {
			count++
		}
	}

	fmt.Fprintln(writer, count)
}
