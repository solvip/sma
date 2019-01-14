package sma_test

import (
	"testing"

	"github.com/solvip/sma"
)

func TestNewMovingAverage(t *testing.T) {
	ma := sma.NewMovingAverage(10)

	if ma.(*sma.SimpleMovingAverage).WindowSize != 10 {
		t.Fatalf("expected WindowSize to be set to 10")
	}
}

func TestValueIsZeroAtStart(t *testing.T) {
	ma := sma.NewMovingAverage(10)
	if val := ma.Value(); val != 0 {
		t.Fatalf("expected Value to be zero when unitialized; instead got %f", val)
	}
}

func TestAdd(t *testing.T) {
	ma := sma.NewMovingAverage(5)

	seq := []struct {
		input    float64
		expected float64
	}{
		{1, 1},
		{2, 1.5},
		{3, 2},
		{4, 2.5},
		{5, 3},
		{6, 4},    // 1 drops off
		{7, 5},    // 2 drops off
		{8, 6},    // 3 drops off
		{9, 7},    // 4 drops off
		{10, 8},   // 5 drops off
		{10, 8.8}, // 6 drops off
		{10, 9.4}, // 7 drops off
		{10, 9.8}, // 8 drops off
	}

	for _, test := range seq {
		ma.Add(test.input)
		if val := ma.Value(); val != test.expected {
			t.Fatalf("epxected ma.Value() == %f, not %f", test.expected, val)
		}
	}
}

func TestSet(t *testing.T) {
	ma := sma.NewMovingAverage(3)

	ma.Set(100.0)
	if ma.Value() != 100.0 {
		t.Fatalf("expected ma.Value() == 100.0; not %f", ma.Value())
	}

	ma.Add(50.0)
	if ma.Value() != 75.0 {
		t.Fatalf("expected ma.Value() == 75.0; not %f", ma.Value())
	}

	ma.Set(0.0)
	if ma.Value() != 0.0 {
		t.Fatalf("expected ma.Value() == 0.0; not %f", ma.Value())
	}

	ma.Add(100.0)
	if ma.Value() != 50.0 {
		t.Fatalf("expected ma.Value() == 50.0; not %f", ma.Value())
	}
}
