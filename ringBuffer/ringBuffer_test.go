package ringBuffer

import (
	"sync"
	"sync/atomic"
	"testing"
)

func consumer(counter *uint64) func([]uint64) {
	mu := &sync.Mutex{}
	return func(items []uint64) {
		mu.Lock()
		for i := range items {
			_ = i
			atomic.AddUint64(counter, 1)
		}
		mu.Unlock()
	}
}

func BenchmarkSyncPool(b *testing.B) {
	received := uint64(0)
	buffer := newSyncPool(consumer(&received))
	total := uint64(0)
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddUint64(&total, 1)
			buffer.Push(1)
		}
	})
	b.Logf("received: %.2f%%\n", (float64(received)/float64(total))*100)
}

func BenchmarkChanDrop(b *testing.B) {
	received := uint64(0)
	buffer := newChanDrop(consumer(&received))
	total := uint64(0)
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddUint64(&total, 1)
			buffer.Push(1)
		}
	})
	b.Logf("received: %.2f%%\n", (float64(received)/float64(total))*100)
}

func BenchmarkMutexLock(b *testing.B) {
	received := uint64(0)
	buffer := newMutexLock(consumer(&received))
	total := uint64(0)
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddUint64(&total, 1)
			buffer.Push(1)
		}
	})
	b.Logf("received: %.2f%%\n", (float64(received)/float64(total))*100)
}

func BenchmarkChanLock(b *testing.B) {
	received := uint64(0)
	buffer := newChanLock(consumer(&received))
	total := uint64(0)
	b.SetBytes(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.AddUint64(&total, 1)
			buffer.Push(1)
		}
	})
	b.Logf("received: %.2f%%\n", (float64(received)/float64(total))*100)
}
