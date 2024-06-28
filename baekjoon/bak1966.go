package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Document struct {
	index    int // 문서의 원래 위치
	priority int // 문서의 중요도
}

func bak1966() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanf(reader, "%d\n", &t)

	for i := 0; i < t; i++ {
		// 문서의 수 n과 우리가 찾고자 하는 문서의 위치 m 입력
		var n, m int
		fmt.Fscanf(reader, "%d %d\n", &n, &m)

		// 문서의 중요도 입력
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		priorities := strings.Split(str, " ")

		// 큐 초기화
		queue := make([]Document, 0)
		for index, value := range priorities {
			priority, _ := strconv.Atoi(value)
			queue = append(queue, Document{index, priority})
		}

		printOrder := 0

		for len(queue) > 0 {
			// 큐의 첫 번째 문서를 꺼내어 가장 높은 우선순위인지 확인
			currentDoc := queue[0]
			queue = queue[1:]

			// 가장 높은 우선순위인지 확인
			isHighest := true
			for _, doc := range queue {
				if doc.priority > currentDoc.priority {
					isHighest = false
					break
				}
			}

			if isHighest {
				// 현재 문서를 출력
				printOrder++
				if currentDoc.index == m {
					fmt.Fprintln(writer, printOrder)
					break
				}
			} else {
				// 우선순위가 높지 않으면 큐의 뒤로 다시 추가
				queue = append(queue, currentDoc)
			}
		}
	}
}
