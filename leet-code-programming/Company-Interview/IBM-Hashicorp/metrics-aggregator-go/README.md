# Metrics Aggregator Challenge

In this problem, you will implement an In-Memory Metrics Aggregator. This system is designed to track execution times (latencies) for various service operations and provide real-time statistical summaries.

## Guidelines

* You have 45 minutes to work on the implementation. Get as far as you can and
  be prepared to discuss your approach and how you would complete the exercise
  afterward.
* Given the time limitation focus on your implementation and mention
  performance or style concerns to the interviewer for later discussion.

### Suggested Approach

* Read through this README to understand the problem.
* Read through `metrics_test.go` and run tests to see where they fail.
* Try to get 1 test passing at a time. From top to bottom in `metrics_test.go`.
 
## Functional Requirements

The aggregator must support three primary operations:

1. **Record(name, durationMs)**: Save a data point with the current timestamp.
2. **GetAverage(name)**: Calculate the mean of non-expired points.
3. **GetSnapshot()**: Return a map of `Stats` (Min, Max, Avg) for all metrics with valid data.


## The Sliding Window Constraint (Critical)

A data point is "expired" if it is older than the `window` duration provided during initialization.

* The MetricsStore is initialized with a windowMs (e.g., 60,000ms for 1 minute).

* Any data point older than windowMs relative to the current system time is considered "expired."

* Expired data must be excluded from getAverage and getSnapshot.


## Instructions
- Use `go test ./...` to run the test suite.
- Focus on correct time-window logic using the `time` package.
- Be prepared to discuss how you would make this safe for concurrent use (goroutines).

## Example Scenario

Assume windowMs = 100 (100 milliseconds):

```
Time (ms),      Action,                 Resulting Average,      Notes
T=0,        "record(""api"", 100)",         100.0,
T=50,       "record(""api"", 200)",         150.0,             Both points are < 100ms old
T=120,      "getAverage(""api"")",          200.0,             The T=0 point has expired
T=250,      "getAverage(""api"")",          0.0,               All points have expired
```
