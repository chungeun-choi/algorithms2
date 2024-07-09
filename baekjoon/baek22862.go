package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Baek22862() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 첫 번째 줄 입력 받기
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	nk := strings.Split(line, " ")
	n, _ := strconv.Atoi(nk[0])
	k, _ := strconv.Atoi(nk[1])

	// 두 번째 줄 입력 받기
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	strA := strings.Split(line, " ")
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i], _ = strconv.Atoi(strA[i])
	}

	end := 0    // 끝 포인터
	answer := 0 // 정답: 가장 긴 짝수로 이루어진 연속 부분 수열 길이
	tmp := 0    // 현재 만들고 있는 가장 긴 짝수로 이루어진 연속 부분 수열 길이
	odd := 0    // 안에 있는 홀수 개수

	for start := 0; start < n; start++ {
		for odd <= k && end < n {
			if A[end]%2 != 0 {
				odd++
			} else {
				tmp++
			}
			end++

			if start == 0 && end == n {
				answer = tmp
				break
			}
		}
		if odd == k+1 {
			if tmp > answer {
				answer = tmp
			}
		}
		if A[start]%2 != 0 {
			odd--
		} else {
			tmp--
		}
	}

	fmt.Fprintln(writer, answer)
}
