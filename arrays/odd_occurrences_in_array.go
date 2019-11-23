package solution

// Solution find odd occurrences of a number in array
func Solution(A []int) int {
	occurrences := make(map[int]int)

	for _, value := range A {
		occurrences[value] = occurrences[value] + 1
	}

	for key, value := range occurrences {
		if value%2 != 0 {
			return key
		}
	}

	return 0
}
