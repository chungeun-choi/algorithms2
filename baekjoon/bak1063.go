package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bak1063() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var (
		kingPos  [2]int
		stonePos [2]int
		moveList []string
		moveDis  map[string][2]int = map[string][2]int{
			"R":  {1, 0},
			"L":  {-1, 0},
			"B":  {0, -1},
			"T":  {0, 1},
			"RT": {1, 1},
			"LT": {-1, 1},
			"RB": {1, -1},
			"LB": {-1, -1},
		}
		n int
	)

	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strs := strings.Split(str, " ")

	// 초기 왕의 위치
	x := rune(strs[0][0])
	y, _ := strconv.Atoi(string(strs[0][1]))
	kingPos = [2]int{int(x - 'A' + 1), y}

	// 초기 돌의 위치
	sx := rune(strs[1][0])
	sy, _ := strconv.Atoi(string(strs[1][1]))
	stonePos = [2]int{int(sx - 'A' + 1), sy}

	n, _ = strconv.Atoi(strs[2])

	moveList = make([]string, n)
	for i := 0; i < n; i++ {
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)

		moveList[i] = str
	}

	for _, moveObj := range moveList {
		cacl, _ := moveDis[moveObj]
		newKingPos := [2]int{kingPos[0] + cacl[0], kingPos[1] + cacl[1]}

		// 새로운 왕의 위치가 체스판을 벗어나지 않는지 확인
		if newKingPos[0] < 1 || newKingPos[0] > 8 || newKingPos[1] < 1 || newKingPos[1] > 8 {
			continue
		}

		if newKingPos == stonePos {
			newStonePos := [2]int{stonePos[0] + cacl[0], stonePos[1] + cacl[1]}
			// 새로운 돌의 위치가 체스판을 벗어나지 않는지 확인
			if newStonePos[0] < 1 || newStonePos[0] > 8 || newStonePos[1] < 1 || newStonePos[1] > 8 {
				continue
			}
			stonePos = newStonePos
		}

		kingPos = newKingPos
	}

	// 결과 출력
	fmt.Fprintf(writer, "%c%d\n%c%d\n", kingPos[0]+'A'-1, kingPos[1], stonePos[0]+'A'-1, stonePos[1])
}
