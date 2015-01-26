package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

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

func init() {
	Plugins = append(Plugins, PdPlugin{})
}

func (p PdPlugin) Help() string {
	return "```Set pagerduty maintenance window. <window time> should have a unit suffix.\nms : Millisecond\ns : Second\nm : Minute\nh : Hour \nUsage: !pd mw <window time>```"
}

func (p PdPlugin) Name() string {
	return "pd"
}

func (p PdPlugin) Execute(command []string) string {
	if len(command) < 2 {
		return p.Help()
	}
	if command[0] != "mw" {
		return "Can only set MW at this time"
	}

	mwDur, err := time.ParseDuration(command[1])
	if err != nil {
		return fmt.Sprintf("Error parsing mw duration. Error: %v", err)
	}
	return setMw(mwDur)
}

func setMw(mwDur time.Duration) string {
	if configMap["pd_token"] == "" || configMap["pd_domain"] == "" {
		return "No padgerduty token or domain set"
	}
	if configMap["pd_services"] == "" {
		return "No services configured!"
	}

	var mw pdMw
	now := time.Now()
	mw.MaintenanceWindow.ServiceIds = strings.Split(configMap["pd_services"], ",")
	mw.MaintenanceWindow.Description = "Gopherbot setting the MW"
	mw.MaintenanceWindow.StartTime = now.Format(time.RFC3339)
	mw.MaintenanceWindow.EndTime = now.Add(mwDur).Format(time.RFC3339)
	mw.RequesterID = configMap["pd_reqid"]

	client := &http.Client{}
	mwJson, err := json.Marshal(mw)
	if err != nil {
		return "Error marshaling mw request"
	}
	req, err := http.NewRequest("POST", "https://"+configMap["pd_domain"]+".pagerduty.com/api/v1/maintenance_windows", bytes.NewReader(mwJson))
	if err != nil {
		return "Error making a new request"
	}

	req.Header.Add("Authorization", "Token token="+configMap["pd_token"])
	req.Header.Add("Content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "Error calling pagerduty API"
	}

	if resp.StatusCode != 201 {
		return fmt.Sprintf("Error posting to pagerduty API. Code: %v", resp.StatusCode)
	}
	return "Set PD Maintenance Window"
}
