package models

// Test A test in a system
type Test struct {
	TestType string            `json:"test_type"`
	URL      string            `json:"url"`
	Headers  map[string]string `json:"headers"`
	Body     map[string]string `json:"body"`
	Alerts   []*Alert          `json:"alerts"`
}
