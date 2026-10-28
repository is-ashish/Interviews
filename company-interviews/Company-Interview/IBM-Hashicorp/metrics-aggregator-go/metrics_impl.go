package metrics

import (
	"time"
)

type dataPoint struct {
	timestamp time.Time
	value     int
}

type metricsStoreImpl struct {
	window time.Duration
	// TODO: Define your internal data structures here (e.g., a map)
}

// NewMetricsStore creates a new instance of the aggregator.
func NewMetricsStore(window time.Duration) MetricsStore {
	return &metricsStoreImpl{
		window: window,
	}
}

type action struct {
	Name string
	Duration int
}


var store (
	data = make(map[int]action)
)

func (s *metricsStoreImpl) Record(name string, durationMs int) {
	// TODO: Implement recording logic
}

func (s *metricsStoreImpl) GetAverage(name string) float64 {
	// TODO: Implement calculation logic with time-based eviction
	return 0.0
}

func (s *metricsStoreImpl) GetSnapshot() map[string]Stats {
	// TODO: Implement snapshot logic for all metrics
	return make(map[string]Stats)
}