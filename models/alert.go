package models

// Alert A alert (warning/error) based on Test response
type Alert struct {
	Variable      string  `json:"variable"`
	Comparison    string  `json:"comparison"`
	ExactValue    float64 `json:"exact_value"`
	IntervalStart float64 `json:"interval_start"`
	IntervalEnd   float64 `json:"interval_end"`
	Mesage        float64 `json:"mesage"`
	IsError       bool    `json:"is_error"`
}
