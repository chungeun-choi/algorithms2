package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func back2493() []int {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Read the number of elements
	var n int
	fmt.Fscanln(reader, &n)

	var towers []int = make([]int, n)
	text, _ := reader.ReadString('\n')
	texts := strings.Split(text, " ")
	for i, tower := range texts {
		tower = strings.TrimSpace(tower)
		towers[i], _ = strconv.Atoi(tower)
	}

	var result []int
	var stack [][2]int

	/*/
	stack 에서 pop 을 할 수 있는 이유는 나보다 큰 값일 때만 수신을 할 수 있기때문임....
	예를 들어 5, 7, 2 값이 존재하고 7의 수신여부를 확인하는 동작을 진행할 때 2의 값은 5값에서 수신 받는게 아닌
	7의 값에서 수신 받기 때문에 5의 값을 pop 할 수 있는 것
	추가적인 예시로 5, 7, 8 일 경우 맨 마지막의 8의 값은 5,7 이 모두 작기 때문에 수신을 받지 못함 또한 이후 작업은
	8이 가장 높음에 따라 먼저 수신할 것임으로 순서는 중요치 않음
	*/

	for i := 0; i < len(towers); i++ {
		for len(stack) > 0 && stack[len(stack)-1][1] <= towers[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i] = stack[len(stack)-1][0]
		}

		stack = append(stack, [2]int{i + 1, towers[i]})

	}

	// Output the result using writer
	for i, res := range result {
		if i > 0 {
			writer.WriteString(" ")
		}
		writer.WriteString(strconv.Itoa(res))
	}
	writer.Flush()

	return result
}
