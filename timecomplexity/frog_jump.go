package timecomplexity

// TimeComplexity calculates frog jumps to Y coordinates
func TimeComplexity(X, Y, D int) int {
	distance := Y - X

	if distance % D == 0 {
		return distance / D
	}

	return (distance / D) + 1
}
