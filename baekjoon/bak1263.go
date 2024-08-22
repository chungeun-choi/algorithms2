package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Task 구조체는 각 작업의 소요 시간(duration)과 마감 시간(deadline)을 포함합니다.
type Task struct {
	duration int
	deadline int
}

func Baek1263() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 첫 줄에서 작업의 수를 읽어옴
	input, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(input))

	tasks := make([]Task, N)

	// 각 작업의 소요 시간과 마감 시간을 읽어옴
	for i := 0; i < N; i++ {
		input, _ = reader.ReadString('\n')
		values := strings.Fields(input)
		duration, _ := strconv.Atoi(values[0])
		deadline, _ := strconv.Atoi(values[1])
		tasks[i] = Task{duration, deadline}
	}

	// 마감 시간을 기준으로 내림차순 정렬
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].deadline > tasks[j].deadline
	})

	// 가장 늦게 시작할 수 있는 시간 계산
	ans := tasks[0].deadline - tasks[0].duration
	for i := 1; i < N; i++ {
		// i번째 일을 끝마쳐야 하는 시간이 현재 시간보다 앞선다면
		if ans > tasks[i].deadline {
			ans = tasks[i].deadline - tasks[i].duration
		} else {
			ans -= tasks[i].duration
		}
	}

	// 결과 출력
	if ans >= 0 {
		fmt.Fprintln(writer, ans)
	} else {
		fmt.Fprintln(writer, -1)
	}
}
