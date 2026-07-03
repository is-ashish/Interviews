package metrics

import (
	"testing"
	"time"
)

func TestBasicRecording(t *testing.T) {
	store := NewMetricsStore(60 * time.Second)
	store.Record("db_query", 100)
	store.Record("db_query", 200)

	avg := store.GetAverage("db_query")
	if avg != 150.0 {
		t.Errorf("Expected average 150.0, got %f", avg)
	}
}

func TestTimeWindowEviction(t *testing.T) {
	// Short window of 200ms
	store := NewMetricsStore(200 * time.Millisecond)

	store.Record("api_latency", 100)
	time.Sleep(100 * time.Millisecond)
	store.Record("api_latency", 200)

	// Both should be valid at this point
	if store.GetAverage("api_latency") != 150.0 {
		t.Errorf("Expected 150.0, got %f", store.GetAverage("api_latency"))
	}

	time.Sleep(150 * time.Millisecond)
	// The 100ms record should have expired
	if store.GetAverage("api_latency") != 200.0 {
		t.Errorf("Expected 200.0 after eviction, got %f", store.GetAverage("api_latency"))
	}
}

func TestSnapshot(t *testing.T) {
	store := NewMetricsStore(150 * time.Millisecond)
	store.Record("metric_a", 10)
	
	time.Sleep(200 * time.Millisecond)
	store.Record("metric_b", 20)

	snapshot := store.GetSnapshot()
	
	if _, ok := snapshot["metric_a"]; ok {
		t.Error("metric_a should have expired from snapshot")
	}
	if stats, ok := snapshot["metric_b"]; !ok || stats.Average != 20.0 {
		t.Error("metric_b missing or incorrect in snapshot")
	}
}