package solution

import (
	"fmt"
	"testing"
)

var number = 1041
var binary = fmt.Sprintf("%b", number)
var expected = 5

func TestBinaryGap(t *testing.T) {
	got := BinaryGap(number)

	if got != expected {
		t.Errorf("Number %d binary of it %s got %d expected %d", number, binary, got, expected)
	}
}

func BenchmarkBinaryGap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BinaryGap(number)
	}
}
