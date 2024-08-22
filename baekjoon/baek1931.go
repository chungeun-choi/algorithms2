package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Baek1931() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 받기
	input, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(input))

	// 시간 간격을 저장할 슬라이스
	arr := make([][2]int, n)

	for i := 0; i < n; i++ {
		input, _ = reader.ReadString('\n')
		values := strings.Fields(input)
		start, _ := strconv.Atoi(values[0])
		end, _ := strconv.Atoi(values[1])
		arr[i] = [2]int{start, end}
	}

	// 끝나는 시간을 기준으로, 끝나는 시간이 같으면 시작 시간을 기준으로 정렬
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] < arr[j][1]
	})

	// 겹치지 않는 최대 간격을 찾음
	endPoint := 0
	answer := 0

	for _, interval := range arr {
		newStart, newEnd := interval[0], interval[1]
		if endPoint <= newStart {
			answer++
			endPoint = newEnd
		}
	}

	fmt.Fprintln(writer, answer)
}
