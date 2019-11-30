package timecomplexity

import (
	"math/rand"
)

func quickSort(A []int) []int {
	if len(A) < 2 {
		return A
	}

	left, right := 0, len(A)-1

	pivot := rand.Int() % len(A)

	A[pivot], A[right] = A[right], A[pivot]

	for i, _ := range A {
		if A[i] < A[right] {
			A[left], A[i] = A[i], A[left]
			left++
		}
	}

	A[left], A[right] = A[right], A[left]

	quickSort(A[:left])
	quickSort(A[left+1:])

	return A
}

// PermMissingElem find the lost value
func PermMissingElem(A []int) int {
	quickSort(A)

	if len(A) == 0 || A[0] != 1 {
		return 1
	}

	for i := 0; i < len(A)-1; i++ {
		if A[i]+1 != A[i+1] {
			return A[i] + 1
		}
	}

	return A[len(A)-1] + 1
}
