package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 이진 검색 함수
func binarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return true
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func bak1920() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	// n 입력 및 nList 읽기
	fmt.Fscanln(reader, &n)
	nList := make([]int, n)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	textList := strings.Split(text, " ")

	for i, _ := range textList {
		nList[i], _ = strconv.Atoi(textList[i])
	}

	var m int
	// m 입력 및 mList 읽기
	fmt.Fscanln(reader, &m)
	mList := make([]int, m)
	text, _ = reader.ReadString('\n')
	text = strings.TrimSpace(text)
	textList = strings.Split(text, " ")

	for i, _ := range textList {
		mList[i], _ = strconv.Atoi(textList[i])
	}

	// nList 정렬
	sort.Slice(nList, func(i, j int) bool {
		return nList[i] < nList[j]
	})

	// mList의 각 요소에 대해 이진 검색 수행 및 결과 출력
	var result strings.Builder
	for i := 0; i < m; i++ {
		if binarySearch(nList, mList[i]) {
			result.WriteString("1\n")
		} else {
			result.WriteString("0\n")
		}
	}

	fmt.Print(result.String())
}
