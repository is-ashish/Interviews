package metrics

// Stats represents the aggregated data for a metric.
type Stats struct {
	Min     int
	Max     int
	Average float64
}

// MetricsStore is the interface the candidate must implement.
type MetricsStore interface {
	// Record tracks a new execution time for a specific metric.
	Record(name string, durationMs int)

	// GetAverage returns the mean duration for a metric within the valid window.
	// Returns 0.0 if no data exists.
	GetAverage(name string) float64

	// GetSnapshot returns a summary for all active metrics.
	// A metric is active if it has at least one non-expired data point.
	GetSnapshot() map[string]Stats
}