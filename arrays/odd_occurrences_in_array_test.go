package solution

import "testing"

var arrayOddOccurrencesInArray = []int{9, 3, 9, 3, 9, 7, 9}
var expectedOddOccurrencesInArray = 7

func TestOddOccurrencesInArray(t *testing.T) {
	gotOddOccurrencesInArray := OddOccurrencesInArray(arrayOddOccurrencesInArray)

	if gotOddOccurrencesInArray != expectedOddOccurrencesInArray {
		t.Errorf("Array is %v, found odd occurrences is %d expected %d", arrayOddOccurrencesInArray, gotOddOccurrencesInArray, expectedOddOccurrencesInArray)
	}
}

func BenchmarkOddOccurrencesInArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		OddOccurrencesInArray(arrayOddOccurrencesInArray)
	}
}
