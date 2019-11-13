package ringBuffer

type chanDrop struct {
	items   chan uint64
	consume func([]uint64)
}

func newChanDrop(consume func([]uint64)) *chanDrop {
	b := &chanDrop{
		items: make(chan uint64, 64),
	}
	go b.process()
	return b
}

func (b *chanDrop) Push(item uint64) {
	select {
	case b.items <- item:
	default:
	}
}

func (b *chanDrop) process() {
	items := make([]uint64, 0, 64)
	for item := range b.items {
		items = append(items, item)
		if len(items) == 64 {
			b.consume(items)
		}
		items = items[:0]
	}
}
