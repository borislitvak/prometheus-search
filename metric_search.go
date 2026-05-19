// Package main provides a Prometheus metric semantic search implementation.
// STUB IMPLEMENTATION - Tests should fail with this stub.
package main

// Metric represents a Prometheus metric with metadata.
type Metric struct {
	Name           string
	HelpText       string
	MetricType     string
	Labels         map[string]string
	DashboardCount int
	QueryCount     int
}

// SearchResult represents a search result with metric and score.
type SearchResult struct {
	Metric Metric
	Score  float64
}

// MetricSearch provides semantic search for Prometheus metrics using CPU embeddings.
type MetricSearch struct {
	metrics []Metric
}

// NewMetricSearch creates a new metric search engine with embeddings.
func NewMetricSearch() *MetricSearch {
	return &MetricSearch{
		metrics: make([]Metric, 0),
	}
}

// AddMetric adds a metric to the search index.
func (ms *MetricSearch) AddMetric(metric Metric) {
	ms.metrics = append(ms.metrics, metric)
}

// AddMetrics adds multiple metrics to the index.
func (ms *MetricSearch) AddMetrics(metrics []Metric) {
	for _, metric := range metrics {
		ms.AddMetric(metric)
	}
}

// Search performs semantic search using CPU embeddings.
// STUB: Returns empty results to make tests fail at baseline.
func (ms *MetricSearch) Search(query string, filters map[string]string, limit int) []SearchResult {
	// Stub implementation - returns empty to make tests fail
	return []SearchResult{}
}

// createTestIndex creates a test index with sample metrics.
func createTestIndex() *MetricSearch {
	search := NewMetricSearch()

	metrics := []Metric{
		{Name: "container_cpu_usage_seconds_total", HelpText: "Cumulative cpu time consumed by the container in core-seconds", DashboardCount: 50, QueryCount: 100},
		{Name: "container_cpu_cfs_throttled_seconds_total"},
		{Name: "container_cpu_cfs_periods_total"},
		{Name: "process_cpu_seconds_total"},
		{Name: "node_cpu_seconds_total"},
		{Name: "cpu_utilization_percent", HelpText: "Percentage of CPU time used"},
		{Name: "cpu_usage_percent", HelpText: "CPU usage as a percentage", DashboardCount: 5, QueryCount: 10},
		{Name: "system_cpu_load_1m"},
		{Name: "cpu_throttle_count"},
		{Name: "container_memory_working_set_bytes", HelpText: "Current working set of the container in bytes"},
		{Name: "container_memory_rss"},
		{Name: "container_memory_cache"},
		{Name: "process_resident_memory_bytes"},
		{Name: "node_memory_MemAvailable_bytes"},
		{Name: "mem_used_percent", HelpText: "Percentage of memory used"},
		{Name: "memory_utilization", HelpText: "Memory utilization percentage"},
		{Name: "http_requests_total", HelpText: "Total number of HTTP requests"},
		{Name: "http_request_duration_seconds", HelpText: "Duration of HTTP requests in seconds"},
		{Name: "http_request_duration_seconds_bucket"},
		{Name: "http_response_size_bytes"},
		{Name: "http_request_size_bytes"},
		{Name: "http_requests_in_flight"},
		{Name: "nginx_http_requests_total"},
		{Name: "http_server_requests_seconds"},
		{Name: "error_count_total", HelpText: "Total number of errors"},
		{Name: "exceptions_total", HelpText: "Total number of exceptions"},
		{Name: "failed_requests_total", HelpText: "Total number of failed requests"},
		{Name: "request_duration_seconds", HelpText: "Request duration in seconds"},
		{Name: "response_time_ms", HelpText: "Response time in milliseconds"},
		{Name: "processing_time_seconds"},
		{Name: "latency_histogram_seconds", HelpText: "Latency histogram"},
		{Name: "api_latency_seconds"},
		{Name: "service_x_cpu_usage", Labels: map[string]string{"service": "X"}},
		{Name: "service_x_memory_used", Labels: map[string]string{"service": "X"}},
		{Name: "service_x_request_count", Labels: map[string]string{"service": "X"}},
		{Name: "service_y_cpu_utilization", Labels: map[string]string{"service": "Y", "app": "service_y"}},
		{Name: "service_y_mem_percent", Labels: map[string]string{"service": "Y", "app": "service_y"}},
	}

	search.AddMetrics(metrics)
	return search
}
