package solution

// CyclicRotation right shift elements positions by K times
func CyclicRotation(A []int, K int) []int {
	if len(A) == 0 || K == len(A) {
		return A
	}

	if K < len(A) {
		K = len(A) - K
	}

	if K > len(A) {
		K = len(A) - (K % len(A))
	}

	var nA = make([]int, len(A))
	copy(nA[0:], A[K:])
	copy(nA[len(A)-K:], A[:K])

	return nA
}
