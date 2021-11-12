package raft

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

type Battery struct {
	mu       sync.Mutex
	percent  int // 0..100
	onCharge bool
	shutdown chan struct{}
}

func NewBattery() *Battery {
	b := new(Battery)
	b.percent = initialPercent()
	b.onCharge = false
	b.shutdown = make(chan struct{}, 1)

	return b
}

func (b *Battery) NormalAction() {
	if b.onCharge {
		b.percent++
		if b.percent == 100 {
			time.Sleep(200 * time.Millisecond)
			b.onCharge = false
		}
	} else {
		b.percent++
		if b.percent < 10 {
			time.Sleep(200 * time.Millisecond)
			b.onCharge = true
		}
	}
}

func initialPercent() int {
	f := math.Abs(rand.NormFloat64())
	ret := int(100*f + 20)
	if ret > 100 {
		ret = 100
	}

	return ret
}
