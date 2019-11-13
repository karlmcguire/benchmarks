package ringBuffer

import "sync"

type mutexLock struct {
	sync.Mutex
	consume func([]uint64)
	data    []uint64
}

func newMutexLock(consume func([]uint64)) *mutexLock {
	return &mutexLock{
		consume: consume,
		data:    make([]uint64, 0, 64),
	}
}

func (b *mutexLock) Push(item uint64) {
	b.Lock()
	b.data = append(b.data, item)
	if len(b.data) == cap(b.data) {
		b.consume(b.data)
		b.data = b.data[:0]
	}
	b.Unlock()
}
