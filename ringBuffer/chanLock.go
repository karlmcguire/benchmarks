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
		items := make([]uint64, 0)
		for len(b.items) > 0 {
			items = append(items, <-b.items)
			if len(items) == 64 {
				break
			}
		}
		b.consume(items)
	}
}
