package ringBuffer

import (
	"sync"
	"testing"
)

func consumer() func([]uint64) {
	count := make(map[uint64]uint64)
	mu := &sync.Mutex{}
	return func(items []uint64) {
		mu.Lock()
		for i := range items {
			count[items[i]]++
		}
		mu.Unlock()
	}
}

func BenchmarkSyncPool(b *testing.B) {
	buffer := newSyncPool(consumer())
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buffer.Push(1)
		}
	})
}

func BenchmarkChanDrop(b *testing.B) {
	buffer := newChanDrop(consumer())
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buffer.Push(1)
		}
	})
}

func BenchmarkMutexLock(b *testing.B) {
	buffer := newMutexLock(consumer())
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buffer.Push(1)
		}
	})
}

func BenchmarkChanLock(b *testing.B) {
	buffer := newChanLock(consumer())
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			buffer.Push(1)
		}
	})
}
