package timecomplexity

import "testing"

var x, y, d = 10, 85, 30
var expected = 3

func TestTimeComplexity(t *testing.T) {
	got := TimeComplexity(x, y, d)

	if got != expected {
		t.Errorf("expected %d got %d\n", expected, got)
	}
}

func BenchmarkTimeComplexity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TimeComplexity(x, y, d)
	}
}