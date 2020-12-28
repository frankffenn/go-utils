package impl

import (
	"math"
	"math/rand"
	"time"
)

type Backoff struct {
	minDelay time.Duration
	maxDelay time.Duration
}

func (b *Backoff) next(attempt int) time.Duration {
	if attempt < 0 {
		return b.minDelay
	}

	minf := float64(b.minDelay)
	durf := minf * math.Pow(1.5, float64(attempt))
	durf = durf + rand.Float64()*minf

	delay := time.Duration(durf)
	if delay > b.maxDelay {
		return b.maxDelay
	}

	return delay
}
