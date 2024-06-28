package programers

func Solution2(clothes [][]string) int {
	var (
		result     int            = 1
		clothesMap map[string]int = make(map[string]int)
	)

	for _, clothe := range clothes {
		if _, ok := clothesMap[clothe[1]]; ok {
			clothesMap[clothe[1]]++
		} else {
			clothesMap[clothe[1]] = 1
		}
	}

	for _, value := range clothesMap {
		result *= value + 1
	}

	return result - 1
}
