package arrays

func Sum(numbers []int) int {
	var length (int) = 0
	for _, number := range numbers {
		length += number
	}
	return length
}

// range help iterate over all the elements in an array
