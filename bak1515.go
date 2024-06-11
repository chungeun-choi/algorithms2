package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bak1515() {
	// 입력을 위한 Reader와 출력을 위한 Writer 준비
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력받은 문자열 읽기
	inputStr, _ := reader.ReadString('\n')
	inputStr = strings.TrimSpace(inputStr)

	// 문자열의 현재 인덱스
	strIndex := 0

	// 자연수 시작
	start := 1

	for strIndex < len(inputStr) {
		// start 숫자를 문자열로 변환
		currentStr := strconv.Itoa(start)

		// currentStr의 각 문자와 inputStr의 현재 문자를 비교
		for _, char := range currentStr {
			// 현재 문자가 inputStr의 현재 인덱스와 일치하는지 확인
			if char == rune(inputStr[strIndex]) {
				// 일치하면 다음 문자로 이동
				strIndex++
				// 모든 문자가 매칭되면 루프 종료
				if strIndex == len(inputStr) {
					break
				}
			}
		}
		// 다음 숫자로 이동
		start++
	}

	// 최종 숫자 출력 (start는 마지막에 1 더해져 있으므로 -1 해줌)
	fmt.Fprintln(writer, start-1)
}
