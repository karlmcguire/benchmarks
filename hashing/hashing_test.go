package hashing

import "testing"

func BenchmarkMemHash(b *testing.B) {
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		MemHash("key")
	}
}

func BenchmarkHighwayHash(b *testing.B) {
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		HighwayHash("key")
	}
}

func BenchmarkXXHash(b *testing.B) {
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		XXHash("key")
	}
}

func BenchmarkFarmHash(b *testing.B) {
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		FarmHash("key")
	}
}

func BenchmarkSipHash(b *testing.B) {
	b.SetBytes(1)
	for n := 0; n < b.N; n++ {
		SipHash("key")
	}
}
