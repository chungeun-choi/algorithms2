package programers

func Solution1(phone_number string) string {
	var result string
	point := len(phone_number) - 4

	for i := 0; i < point; i++ {
		result += "*"
	}
	result += phone_number[point:]

	return result
}
