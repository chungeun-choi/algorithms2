package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func baek1092() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var (
		n, m          int
		cranes, boxes []int
	)

	// 크레인 수 입력
	fmt.Fscanln(reader, &n)
	// 크레인 용량 입력
	craneInput, _ := reader.ReadString('\n')
	craneInput = strings.TrimSpace(craneInput)
	craneStrings := strings.Split(craneInput, " ")
	cranes = make([]int, n)
	for i, v := range craneStrings {
		cranes[i], _ = strconv.Atoi(v)
	}

	// 박스 수 입력
	fmt.Fscanln(reader, &m)
	// 박스 무게 입력
	boxInput, _ := reader.ReadString('\n')
	boxInput = strings.TrimSpace(boxInput)
	boxStrings := strings.Split(boxInput, " ")
	boxes = make([]int, m)
	for i, v := range boxStrings {
		boxes[i], _ = strconv.Atoi(v)
	}

	// 크레인과 박스를 내림차순으로 정렬
	sort.Sort(sort.Reverse(sort.IntSlice(cranes)))
	sort.Sort(sort.Reverse(sort.IntSlice(boxes)))

	// 만약 가장 큰 크레인이 가장 무거운 박스를 들 수 없으면 불가능
	if len(boxes) > 0 && boxes[0] > cranes[0] {
		fmt.Fprintln(writer, -1)
		return
	}

	// 라운드 수 계산
	cnt := 0

	for len(boxes) > 0 {
		cnt++
		// 각 크레인에 대해 가능한 가장 무거운 박스를 찾기
		craneIndex := 0
		remainingBoxes := []int{}

		for _, box := range boxes {
			if craneIndex < n && cranes[craneIndex] >= box {
				// 크레인이 이 박스를 옮길 수 있는 경우
				craneIndex++
			} else {
				// 이번 라운드에 옮길 수 없는 박스는 다음 라운드로 남겨둠
				remainingBoxes = append(remainingBoxes, box)
			}
		}

		// 남은 박스로 boxes 슬라이스를 갱신
		boxes = remainingBoxes
	}

	fmt.Fprintln(writer, cnt)
}
