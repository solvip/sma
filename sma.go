// Package sma implements simple, unweighted moving averages
package sma

// Consistent with VividCortex's ewma package
type MovingAverage interface {
	Add(float64)
	Value() float64
	Set(float64)
}

// NewMovingAverage - return an SMA having a window size of windowSize
func NewMovingAverage(windowSize int) MovingAverage {
	if windowSize < 1 {
		panic("WindowSize must be at least 1")
	}
	return &SimpleMovingAverage{WindowSize: windowSize}
}

type SimpleMovingAverage struct {
	WindowSize int

	samples  []float64
	evictIdx int
	sum      float64
}

func (sma *SimpleMovingAverage) Add(val float64) {
	sma.sum += val

	// If we don't yet have a full window, we simply add
	// and do no evictions
	if len(sma.samples) < sma.WindowSize {
		sma.samples = append(sma.samples, val)
		return
	}

	sma.sum -= sma.samples[sma.evictIdx]
	sma.samples[sma.evictIdx] = val
	sma.evictIdx = (sma.evictIdx + 1) % sma.WindowSize
}

func (sma *SimpleMovingAverage) Value() float64 {
	if len(sma.samples) == 0 {
		return 0
	}

	return sma.sum / float64(len(sma.samples))
}

func (sma *SimpleMovingAverage) Set(val float64) {
	if sma.samples != nil {
		sma.samples = sma.samples[0:0]
		sma.sum = 0
		sma.evictIdx = 0
	}
	sma.Add(val)

	return
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
