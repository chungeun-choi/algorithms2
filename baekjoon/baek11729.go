package baekjoon

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func hanoiTower(n int, start int, end int, writer *bufio.Writer) {
	if n == 1 {
		fmt.Fprintln(writer, start, end)
		return
	}

	hanoiTower(n-1, start, 6-start-end, writer) // 1단계
	fmt.Fprintln(writer, start, end)            // 2단계
	hanoiTower(n-1, 6-start-end, end, writer)   // 3단계
}

func baek11729() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	input, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(input[:len(input)-1])

	// 2^n - 1을 출력
	fmt.Fprintln(writer, int(math.Pow(2, float64(n))-1))

	// 하노이 탑 이동 과정을 출력
	hanoiTower(n, 1, 3, writer)
}
