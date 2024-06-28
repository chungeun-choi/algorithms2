package baekjoon

import (
	"bufio"
	"fmt"
	"os"
)

func bak2193() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	if n == 1 {
		// 길이 1인 이친수는 '1' 하나뿐
		fmt.Fprintln(writer, 1)
		return
	}

	// DP 테이블 초기화
	dp := make([][2]int, n+1)
	dp[1][0] = 0
	dp[1][1] = 1

	// DP 점화식 적용
	for i := 2; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + dp[i-1][1]
		dp[i][1] = dp[i-1][0]
	}

	// 길이 N인 이친수의 개수 출력
	result := dp[n][0] + dp[n][1]
	fmt.Fprintln(writer, result)
}
