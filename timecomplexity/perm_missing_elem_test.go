package timecomplexity

import "testing"

var aPermMissingElem = []int{2, 3, 1, 5}
var ePermMissingElem = 4

func TestPermMissingElem(t *testing.T) {
	got := PermMissingElem(aPermMissingElem)

	if got != ePermMissingElem {
		t.Errorf("expected %d got %d\n", ePermMissingElem, got)
	}
}

func BenchmarkPermMissingElem(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PermMissingElem(aPermMissingElem)
	}
}