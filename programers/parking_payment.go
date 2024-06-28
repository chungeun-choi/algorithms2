package programers

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func convertHourToMinute(s string) int {
	times := strings.Split(s, ":")
	hour, _ := strconv.Atoi(times[0])
	minute, _ := strconv.Atoi(times[1])
	return hour*60 + minute
}

type Data struct {
	time   int
	action string
}

func Solution3(fees []int, records []string) []int {
	var (
		result     []int
		parkingMap map[int]int = make(map[int]int) // carNumber : total parked time
		inTimeMap  map[int]int = make(map[int]int) // carNumber : in time
	)

	// 주차요금에 대한 정책 변수
	var (
		defaultTime int = fees[0]
		defaultPay  int = fees[1]
		perTime     int = fees[2]
		perPay      int = fees[3]
	)

	for _, record := range records {
		datas := strings.Split(record, " ")
		actionTime := convertHourToMinute(datas[0])
		carNumber, _ := strconv.Atoi(datas[1])
		action := datas[2]

		if action == "IN" {
			inTimeMap[carNumber] = actionTime
		} else if action == "OUT" {
			parkingTime := actionTime - inTimeMap[carNumber]
			parkingMap[carNumber] += parkingTime
			delete(inTimeMap, carNumber)
		}
	}

	// 남아 있는 차량 처리 (23:59에 출차 처리)
	endOfDay := convertHourToMinute("23:59")
	for carNumber, inTime := range inTimeMap {
		parkingTime := endOfDay - inTime
		parkingMap[carNumber] += parkingTime
	}

	// 차량 번호 순으로 결과 정렬
	carNumbers := make([]int, 0, len(parkingMap))
	for carNumber := range parkingMap {
		carNumbers = append(carNumbers, carNumber)
	}
	sort.Ints(carNumbers)

	for _, carNumber := range carNumbers {
		totalTime := parkingMap[carNumber]
		if totalTime <= defaultTime {
			result = append(result, defaultPay)
		} else {
			extraTime := totalTime - defaultTime
			extraPay := int(math.Ceil(float64(extraTime)/float64(perTime))) * perPay
			result = append(result, defaultPay+extraPay)
		}
	}

	return result
}
