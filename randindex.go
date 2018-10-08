package main

import (
	"encoding/binary"
	"math"
	"math/rand"
	"time"
)

// RandIndex generates random integers using simple discard method.
type RandIndex struct {
	r *rand.Rand
}

// NewRandIndex returns new random generator.
func NewRandIndex() *RandIndex {
	now := time.Now()
	return &RandIndex{r: rand.New(rand.NewSource(now.Unix()))}
}

// RandInt32 return random 32-bit integer 0 <= n <= max.
func (r *RandIndex) RandInt32(max uint32) uint32 {
	bitsNum := math.Ceil(math.Log2(float64(max)))
	bytesNum := int(math.Ceil(bitsNum / 8.0))
	if bytesNum > 4 {
		bytesNum = 4
	}
	p := make([]byte, bytesNum)
	bytesPad := 4 - bytesNum
	pad := make([]byte, bytesPad)
	for {
		r.r.Read(p)
		num := binary.BigEndian.Uint32(append(pad, p...))
		if num <= max {
			return num
		}
	}
}
