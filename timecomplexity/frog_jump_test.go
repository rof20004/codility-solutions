package timecomplexity

import "testing"

var xTimeComplexity, yTimeComplexity, dTimeComplexity = 10, 85, 30
var expectedTimeComplexity = 3

func TestTimeComplexity(t *testing.T) {
	got := TimeComplexity(xTimeComplexity, yTimeComplexity, dTimeComplexity)

	if got != expectedTimeComplexity {
		t.Errorf("expected %d got %d\n", expectedTimeComplexity, got)
	}
}

func BenchmarkTimeComplexity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TimeComplexity(xTimeComplexity, yTimeComplexity, dTimeComplexity)
	}
}