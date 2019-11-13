package falseSharing

import "sync/atomic"

type Counters interface {
	Increment(int)
}

type Shared struct {
	counters []uint64
}

func NewShared() *Shared {
	return &Shared{
		counters: make([]uint64, 3),
	}
}

func (c *Shared) Increment() {
	atomic.AddUint64(&c.counters[0], 1)
	atomic.AddUint64(&c.counters[1], 1)
	atomic.AddUint64(&c.counters[2], 1)
}

type Padded struct {
	counters [][8]uint64
}

func NewPadded() *Padded {
	return &Padded{
		counters: make([][8]uint64, 3),
	}
}

func (c *Padded) Increment() {
	atomic.AddUint64(&c.counters[0][0], 1)
	atomic.AddUint64(&c.counters[1][0], 1)
	atomic.AddUint64(&c.counters[2][0], 1)
}
