package util

// MovingMedian provides the median value over a moving window.
//
// This type is not concurrency safe.
type MovingMedian struct {
	values []float64
	sorted []float64
	index  int
	size   int
}

func NewMovingMedian(size int) MovingMedian { _ = "STUB: not implemented"; return *new(MovingMedian) }

// Add adds a value to the window, sorts the values, and returns the current median.
func (m *MovingMedian) Add(value float64) float64 { _ = "STUB: not implemented"; return 0 }

// Median returns the current median, else 0 if the window isn't full yet.
func (m *MovingMedian) Median() float64 { _ = "STUB: not implemented"; return 0 }

// Reset resets the window to its initial value.
func (m *MovingMedian) Reset() { _ = "STUB: not implemented"; return }
