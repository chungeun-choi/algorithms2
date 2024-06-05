package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

// 시리얼 번호를 정렬하는 데 사용될 커스텀 타입
type SerialNumbers []string

// 시리얼 번호의 길이로 정렬
func (s SerialNumbers) Len() int {
	return len(s)
}

func (s SerialNumbers) Less(i, j int) bool {
	// 1. 길이가 짧은 것이 먼저 온다.
	if len(s[i]) != len(s[j]) {
		return len(s[i]) < len(s[j])
	}

	// 2. 길이가 같으면, 숫자의 합이 작은 것이 먼저 온다.
	sumI, sumJ := 0, 0
	for _, char := range s[i] {
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			sumI += num
		}
	}
	for _, char := range s[j] {
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			sumJ += num
		}
	}
	if sumI != sumJ {
		return sumI < sumJ
	}

	// 3. 길이와 숫자의 합이 같으면, 사전 순으로 먼저 온다.
	return s[i] < s[j]
}

func (s SerialNumbers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func bak1431() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	serials := make(SerialNumbers, n)
	for i := 0; i < n; i++ {
		serial, _ := reader.ReadString('\n')
		serials[i] = serial[:len(serial)-1] // 개행 문자 제거
	}

	sort.Sort(serials)

	for _, serial := range serials {
		fmt.Fprintln(writer, serial)
	}
}
