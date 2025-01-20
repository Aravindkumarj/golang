package arrays

func Sum(numbers [5]int) int {
	var length (int) = 0
	for i := 0; i < len(numbers); i++ {
		length += numbers[i]
	}
	return length
}
