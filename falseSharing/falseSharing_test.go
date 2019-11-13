package falseSharing

import (
	"testing"
)

func BenchmarkShared(b *testing.B) {
	shared := NewShared(256)
	for n := 0; n < b.N; n++ {
		shared.Increment(n & 255)
	}
}

func BenchmarkPadded(b *testing.B) {
	padded := NewPadded(256)
	for n := 0; n < b.N; n++ {
		padded.Increment(n & 255)
	}
}
