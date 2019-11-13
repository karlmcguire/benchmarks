package falseSharing

import "sync/atomic"

type Counters interface {
	Increment(int)
}

type Shared struct {
	counters []uint64
}

func NewShared(size int) *Shared {
	return &Shared{
		counters: make([]uint64, size),
	}
}

func (c *Shared) Increment(i int) {
	atomic.AddUint64(&c.counters[i], 1)
}

type Padded struct {
	counters []PaddedCounter
}

type PaddedCounter struct {
	counter uint64
	_pad    [8]uint64
}

func NewPadded(size int) *Padded {
	return &Padded{
		counters: make([]PaddedCounter, size),
	}
}

func (c *Padded) Increment(i int) {
	atomic.AddUint64(&c.counters[i].counter, 1)
}
