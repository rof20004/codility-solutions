package timecomplexity

import "testing"

var xFrogJump, yFrogJump, dFrogJump = 10, 85, 30
var eFrogJump = 3

func TestFrogJump(t *testing.T) {
	got := FrogJump(xFrogJump, yFrogJump, dFrogJump)

	if got != eFrogJump {
		t.Errorf("expected %d got %d\n", eFrogJump, got)
	}
}

func BenchmarkFrogJump(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FrogJump(xFrogJump, yFrogJump, dFrogJump)
	}
}