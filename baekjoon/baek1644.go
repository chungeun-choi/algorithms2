package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func Baek1644() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 한 줄 입력을 읽고 정수로 변환
	input, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(input))

	if n < 2 {
		fmt.Fprintln(writer, 0)
		return
	}

	// 주어진 N 이하의 모든 소수를 찾기
	primes := []int{}
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}

	// 두 포인터를 사용하여 연속된 소수의 합 계산
	count := 0
	start, end, sum := 0, 0, 0
	for end <= len(primes) {
		if sum < n {
			if end == len(primes) {
				break
			}
			sum += primes[end]
			end++
		} else if sum > n {
			sum -= primes[start]
			start++
		} else {
			count++
			sum -= primes[start]
			start++
		}
	}

	fmt.Fprintln(writer, count)
}
