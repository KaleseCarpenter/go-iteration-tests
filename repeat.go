package iteration

import "testing"

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		//+= "the Add AND assignment operator", adds the right operand to the left operand & assigns the result to left operand
		repeated += character
	}
	return repeated
}

// run go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	// runs b.N times and measures how long it takes
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
