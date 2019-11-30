package timecomplexity

// FrogJump calculates frog jumps to Y coordinates
func FrogJump(X, Y, D int) int {
	distance := Y - X

	if distance % D == 0 {
		return distance / D
	}

	return (distance / D) + 1
}
