package yapip

import (
	"testing"
)

func TestInExtent(t *testing.T) {
	test := []struct {
		ring     [][]float64
		p        []float64
		expected bool
	}{
		// inside
		{
			ring:     [][]float64{{0, 0}, {0, 1}, {1, 1}, {1, 0}, {0, 0}},
			p:        []float64{0.5, 0.5},
			expected: true,
		},
		// outside
		{
			ring:     [][]float64{{0, 0}, {0, 1}, {1, 1}, {1, 0}, {0, 0}},
			p:        []float64{2, 2},
			expected: false,
		},
		// touches
		{
			ring:     [][]float64{{0, 0}, {0, 1}, {1, 1}, {1, 0}, {0, 0}},
			p:        []float64{0, 0},
			expected: true,
		},
	}

	for _, test := range test {
		got := inExtent(test.p, test.ring)

		if got != test.expected {
			t.Errorf("expected %+v, got: %+v", test.expected, got)
		}
	}
}
