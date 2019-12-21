package models

// Alert A alert (warning/error) based on the Test response
type Alert struct {
	Variable      string
	Comparision   string
	ExactValue    float64
	IntervalStart float64
	IntervalEnd   float64
	Mesage        float64
}
