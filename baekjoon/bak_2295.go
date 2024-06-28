package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func bak2295() {
	// 빠른 입출력을 위한 스캐너와 출력 버퍼 생성
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 수의 개수 N 읽기
	scanner.Scan()
	var N int
	fmt.Sscan(scanner.Text(), &N)

	// 수를 담을 슬라이스 생성
	numbers := make([]int, N)

	// N개의 수 입력받기
	for i := 0; i < N; i++ {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &numbers[i])
	}

	// 숫자들을 오름차순으로 정렬
	sort.Ints(numbers)

	// 두 수의 합으로 만들 수 있는 모든 경우의 수를 저장할 맵
	twoSumMap := make(map[int]bool)

	// 두 수의 합으로 만들 수 있는 모든 경우의 수를 계산
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			sum := numbers[i] + numbers[j]
			twoSumMap[sum] = true
		}
	}

	// 가능한 값 중 가장 큰 값을 찾기
	maxValue := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			diff := numbers[i] - numbers[j]
			if twoSumMap[diff] && numbers[i] > maxValue {
				maxValue = numbers[i]
			}
		}
	}

	// 결과 출력
	fmt.Fprintln(writer, maxValue)
}
