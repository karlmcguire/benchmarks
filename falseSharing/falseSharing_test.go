package falseSharing

import (
	"testing"
)

func BenchmarkShared(b *testing.B) {
	shared := NewShared()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			shared.Increment()
		}
	})
}

func BenchmarkPadded(b *testing.B) {
	padded := NewPadded()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			padded.Increment()
		}
	})
}
