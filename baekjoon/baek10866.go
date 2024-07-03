package baekjoon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Deque 구조체 정의
type Deque struct {
	data []int
}

// PushFront 덱의 앞에 원소를 추가
func (d *Deque) PushFront(x int) {
	d.data = append([]int{x}, d.data...)
}

// PushBack 덱의 뒤에 원소를 추가
func (d *Deque) PushBack(x int) {
	d.data = append(d.data, x)
}

// PopFront 덱의 앞에서 원소를 제거하고 반환
func (d *Deque) PopFront() int {
	if len(d.data) == 0 {
		return -1
	}
	val := d.data[0]
	d.data = d.data[1:]
	return val
}

// PopBack 덱의 뒤에서 원소를 제거하고 반환
func (d *Deque) PopBack() int {
	if len(d.data) == 0 {
		return -1
	}
	val := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]
	return val
}

// Size 덱의 크기 반환
func (d *Deque) Size() int {
	return len(d.data)
}

// Empty 덱이 비었는지 여부 반환
func (d *Deque) Empty() int {
	if len(d.data) == 0 {
		return 1
	}
	return 0
}

// Front 덱의 앞 원소 반환
func (d *Deque) Front() int {
	if len(d.data) == 0 {
		return -1
	}
	return d.data[0]
}

// Back 덱의 뒤 원소 반환
func (d *Deque) Back() int {
	if len(d.data) == 0 {
		return -1
	}
	return d.data[len(d.data)-1]
}

func Baek10866() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	deque := Deque{}

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")

		switch parts[0] {
		case "push_front":
			val, _ := strconv.Atoi(parts[1])
			deque.PushFront(val)
		case "push_back":
			val, _ := strconv.Atoi(parts[1])
			deque.PushBack(val)
		case "pop_front":
			fmt.Fprintln(writer, deque.PopFront())
		case "pop_back":
			fmt.Fprintln(writer, deque.PopBack())
		case "size":
			fmt.Fprintln(writer, deque.Size())
		case "empty":
			fmt.Fprintln(writer, deque.Empty())
		case "front":
			fmt.Fprintln(writer, deque.Front())
		case "back":
			fmt.Fprintln(writer, deque.Back())
		}
	}
}
