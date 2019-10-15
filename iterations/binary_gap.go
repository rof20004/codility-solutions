package solution

import "fmt"

const numberOneBinaryString = "1"

// Solution return maxGap of a number in a binary representation
func Solution(N int) int {
	binary := convertIntToBinaryString(N)
	maxGap := getMaxGapBinaryString(binary)

	return maxGap
}

func convertIntToBinaryString(N int) string {
	return fmt.Sprintf("%b", N)
}

func getMaxGapBinaryString(binary string) (gap int) {
	var numberOneBinaryStringPositions = findNumberOneBinaryStringPositions(binary)

	gap = findMaxGap(numberOneBinaryStringPositions, len(numberOneBinaryStringPositions))

	return gap
}

func findNumberOneBinaryStringPositions(binary string) (numberOneBinaryStringPositions []int) {
	for position, runeValue := range binary {
		if string(runeValue) == numberOneBinaryString {
			numberOneBinaryStringPositions = append(numberOneBinaryStringPositions, position)
		}
	}
	return
}

func findMaxGap(positionsSlice []int, positionsSliceLength int) (maxGap int) {
	for i := 1; i < positionsSliceLength; i++ {
		max := (positionsSlice[i] - positionsSlice[i-1]) - 1

		if max > maxGap {
			maxGap = max
		}
	}
	return
}
