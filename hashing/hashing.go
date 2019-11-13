package hashing

import (
	"unsafe"

	"github.com/cespare/xxhash"
	"github.com/dchest/siphash"
	farm "github.com/dgryski/go-farm"
	"github.com/minio/highwayhash"
)

type stringStruct struct {
	str unsafe.Pointer
	len int
}

//go:noescape
//go:linkname memhash runtime.memhash
func memhash(p unsafe.Pointer, h, s uintptr) uintptr

func MemHash(str string) uint64 {
	ss := (*stringStruct)(unsafe.Pointer(&str))
	return uint64(memhash(ss.str, 0, uintptr(ss.len)))
}

func HighwayHash(str string) uint64 {
	return highwayhash.Sum64([]byte(str), []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	})
}

func XXHash(str string) uint64 {
	return xxhash.Sum64([]byte(str))
}

func FarmHash(str string) uint64 {
	return farm.Hash64([]byte(str))
}

func SipHash(str string) uint64 {
	return siphash.Hash(0, 0, []byte(str))
}
