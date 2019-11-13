package ringBuffer

type chanLock struct {
	items   chan uint64
	consume func([]uint64)
}

func newChanLock(consume func([]uint64)) *chanLock {
	return &chanLock{
		items:   make(chan uint64, 64),
		consume: consume,
	}
}

func (b *chanLock) Push(item uint64) {
	select {
	case b.items <- item:
	default:
		items := make([]uint64, 64)
		for len(b.items) > 0 {
			items = append(items, <-b.items)
		}
		b.consume(items)
	}
}
