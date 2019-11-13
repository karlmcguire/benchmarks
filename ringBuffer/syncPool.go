package ringBuffer

import "sync"

type syncPool struct {
	pool *sync.Pool
}

func newSyncPool(consume func([]uint64)) *syncPool {
	return &syncPool{
		pool: &sync.Pool{
			New: func() interface{} { return newSyncPoolStripe(consume) },
		},
	}
}

func (b *syncPool) Push(item uint64) {
	stripe := b.pool.Get().(*syncPoolStripe)
	stripe.push(item)
	b.pool.Put(stripe)
}

type syncPoolStripe struct {
	consume func([]uint64)
	data    []uint64
}

func newSyncPoolStripe(consume func([]uint64)) *syncPoolStripe {
	return &syncPoolStripe{
		consume: consume,
		data:    make([]uint64, 0, 64),
	}
}

func (s *syncPoolStripe) push(item uint64) {
	s.data = append(s.data, item)
	if len(s.data) >= cap(s.data) {
		s.consume(s.data)
		s.data = s.data[:0]
	}
}
