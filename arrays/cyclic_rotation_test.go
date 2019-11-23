package solution

import (
	"reflect"
	"testing"
)

var arrayCyclicRotation = []int{3, 8, 9, 7, 6}
var expectedCyclicRotation = []int{9, 7, 6, 3, 8}
var timesCyclicRotation = 3

func TestCyclicRotation(t *testing.T) {
	gotCyclicRotation := CyclicRotation(arrayCyclicRotation, timesCyclicRotation)

	if !reflect.DeepEqual(gotCyclicRotation, expectedCyclicRotation) {
		t.Errorf("Array is %v, found odd occurrences is %d expected %d", arrayCyclicRotation, gotCyclicRotation, expectedCyclicRotation)
	}
}

func BenchmarkCyclicRotation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CyclicRotation(arrayCyclicRotation, timesCyclicRotation)
	}
}