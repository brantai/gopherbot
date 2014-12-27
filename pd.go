package main

type PdPlugin struct{}

type pdMw struct {
	MaintenanceWindow struct {
		Description string   `json:"description"`
		EndTime     string   `json:"end_time"`
		ServiceIds  []string `json:"service_ids"`
		StartTime   string   `json:"start_time"`
	} `json:"maintenance_window"`
	RequesterID string `json:"requester_id"`
}
