package solution

import "testing"

var array = []int{9, 3, 9, 3, 9, 7, 9}
var expected = 7

func TestSolution(t *testing.T) {
	got := Solution(array)

	if got != expected {
		t.Errorf("Array is %v, found odd occurrences is %d expected %d", array, got, expected)
	}
}

func BenchmarkSolution(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solution(array)
	}
}
