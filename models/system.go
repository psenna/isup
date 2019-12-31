package models

// System A system that is monitored
type System struct {
	Name     string `json:"name"`
	Group    string `json:"group"`
	Priority uint   `json:"priority"`
	Tests    []Test `json:"tests"`
}
