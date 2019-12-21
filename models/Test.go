package models

// Test A test in a system
type Test struct {
	TestTipe string
	URL      string
	Headers  map[string]string
	body     map[string]string
	Alerts   []*Alert
}
