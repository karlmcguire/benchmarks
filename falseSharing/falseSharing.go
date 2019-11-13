package falseSharing

import "sync/atomic"

type Counters interface {
	Increment(int)
}

type Shared struct {
	c1 uint64
	c2 uint64
	c3 uint64
}

func NewShared() *Shared {
	return &Shared{}
}

func (c *Shared) Increment() {
	atomic.AddUint64(&c.c1, 1)
	atomic.AddUint64(&c.c2, 1)
	atomic.AddUint64(&c.c3, 1)
}

type Padded struct {
	c1  uint64
	_p1 [8]uint64
	c2  uint64
	_p2 [8]uint64
	c3  uint64
	_p3 [8]uint64
}

func NewPadded() *Padded {
	return &Padded{}
}

func (c *Padded) Increment() {
	atomic.AddUint64(&c.c1, 1)
	atomic.AddUint64(&c.c2, 1)
	atomic.AddUint64(&c.c3, 1)
}
