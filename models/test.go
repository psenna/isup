package models

// Test A test in a system
type Test struct {
	TestType       string                 `json:"test_type"`
	URL            string                 `json:"url"`
	Headers        map[string]interface{} `json:"headers"`
	FormParameters map[string]interface{} `json:"form_parameters"`
	Alerts         []*Alert               `json:"alerts"`
}
