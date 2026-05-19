package main

import (
	"testing"
)

func TestTypoTolerance(t *testing.T) {
	search := createTestIndex()

	tests := []struct {
		query    string
		expected []string
	}{
		{"cpu utlization", []string{"cpu_utilization_percent", "cpu_usage_percent"}},
		{"memmory usage", []string{"container_memory_working_set_bytes", "mem_used_percent"}},
		{"http reqests", []string{"http_requests_total", "http_request_duration_seconds"}},
		{"latncy", []string{"latency_histogram_seconds", "api_latency_seconds"}},
	}

	for _, tt := range tests {
		results := search.Search(tt.query, nil, 10)
		resultNames := make([]string, len(results))
		for i, r := range results {
			resultNames[i] = r.Metric.Name
		}

		found := 0
		for _, exp := range tt.expected {
			for _, name := range resultNames {
				if name == exp {
					found++
					break
				}
			}
		}

		if found == 0 {
			t.Errorf("Query %q: expected at least one of %v in top 10, got %v", tt.query, tt.expected, resultNames)
		}
	}
}

func TestAbbreviations(t *testing.T) {
	search := createTestIndex()

	tests := []struct {
		query    string
		expected []string
	}{
		{"cpu util", []string{"cpu_utilization_percent", "cpu_usage_percent"}},
		{"mem usage", []string{"mem_used_percent", "container_memory_working_set_bytes"}},
		{"req rate", []string{"http_requests_total", "service_x_request_count"}},
		{"http req", []string{"http_requests_total", "http_request_duration_seconds"}},
	}

	for _, tt := range tests {
		results := search.Search(tt.query, nil, 10)
		resultNames := make([]string, len(results))
		for i, r := range results {
			resultNames[i] = r.Metric.Name
		}

		found := 0
		for _, exp := range tt.expected {
			for _, name := range resultNames {
				if name == exp {
					found++
					break
				}
			}
		}

		if found == 0 {
			t.Errorf("Query %q: expected at least one of %v, got %v", tt.query, tt.expected, resultNames)
		}
	}
}

func TestServiceFiltering(t *testing.T) {
	search := createTestIndex()

	tests := []struct {
		query    string
		filters  map[string]string
		expected string
	}{
		{"cpu utilization", map[string]string{"service": "X"}, "service_x_cpu_usage"},
		{"memory", map[string]string{"app": "service_y"}, "service_y_mem_percent"},
		{"request count", map[string]string{"service": "X"}, "service_x_request_count"},
	}

	for _, tt := range tests {
		results := search.Search(tt.query, tt.filters, 10)
		if len(results) == 0 {
			t.Errorf("Query %q with filters %v: got no results", tt.query, tt.filters)
			continue
		}

		resultNames := make([]string, len(results))
		for i, r := range results {
			resultNames[i] = r.Metric.Name
		}

		found := false
		for i, name := range resultNames {
			if name == tt.expected {
				found = true
				if i >= 3 {
					t.Errorf("Query %q: expected %q in top 3, got position %d", tt.query, tt.expected, i)
				}
				break
			}
		}

		if !found {
			t.Errorf("Query %q: expected %q in results, got %v", tt.query, tt.expected, resultNames)
		}

		// Verify all results match filters
		for _, r := range results {
			for key, value := range tt.filters {
				if r.Metric.Labels == nil || r.Metric.Labels[key] != value {
					t.Errorf("Query %q: result %q does not match filter %s=%s", tt.query, r.Metric.Name, key, value)
				}
			}
		}
	}
}

func TestEdgeCases(t *testing.T) {
	search := createTestIndex()

	// Empty query
	results := search.Search("", nil, 5)
	if len(results) == 0 {
		t.Error("Empty query should return results")
	}

	// Exact match
	results = search.Search("container_cpu_usage_seconds_total", nil, 5)
	if len(results) == 0 {
		t.Fatal("Exact match query returned no results")
	}
	if results[0].Metric.Name != "container_cpu_usage_seconds_total" {
		t.Errorf("Exact match: expected container_cpu_usage_seconds_total as #1, got %s", results[0].Metric.Name)
	}
	if results[0].Score < 0.99 {
		t.Errorf("Exact match: expected score >= 0.99, got %.3f", results[0].Score)
	}

	// Non-existent metric
	results = search.Search("xyz_nonexistent_metric", nil, 5)
	if len(results) > 0 && results[0].Score >= 0.2 {
		t.Errorf("Non-existent query: expected max score < 0.2, got %.3f", results[0].Score)
	}

	// Single char
	results = search.Search("a", nil, 10)
	// Should not crash
	_ = results
}
