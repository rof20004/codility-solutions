package solution

import "fmt"

// BinaryGap return maxGap of a number in a binary representation
func BinaryGap(N int) int {
	binary := fmt.Sprintf("%b", N)

	var occurrencesOne = make([]int, 0)
	for position, runeValue := range binary {
		if string(runeValue) == "1" {
			occurrencesOne = append(occurrencesOne, position)
		}
	}

	var maxGap int
	for i := 1; i < len(occurrencesOne); i++ {
		max := (occurrencesOne[i] - occurrencesOne[i-1]) - 1

		if max > maxGap {
			maxGap = max
		}
	}

	return maxGap
}
