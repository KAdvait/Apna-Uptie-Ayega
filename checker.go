package main

import (
	"fmt"
	"net/http"
	"time"
)

type EndpointStatus struct {
	URL          string
	StatusCode   int
	LastChecked  time.Time
	HealthySince time.Time
}

var endpoints = []string{
	"https://dev6.tracknmanage.com/AccessPanelAPI/swagger/index.html",
	"https://pgadmin.tracknmanage.com",
}

var statuses = map[string]*EndpointStatus{}

func checkEndpoint(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error checking %s: %v\n", url, err)
		sendSlackNotification(url, "Error: "+err.Error())
		return
	}
	defer resp.Body.Close()
	currentTime := time.Now().Truncate(time.Second)
	status := statuses[url]
	if resp.StatusCode == 200 {
		if status == nil || status.StatusCode != 200 {
			statuses[url] = &EndpointStatus{
				URL:          url,
				StatusCode:   200,
				LastChecked:  currentTime,
				HealthySince: currentTime,
			}
		} else {
			status.LastChecked = currentTime
		}
	} else {
		if status.StatusCode == 200 {
			sendSlackNotification(url, fmt.Sprintf("Non-200 response: %d", resp.StatusCode))
		}
		statuses[url] = &EndpointStatus{
			URL:         url,
			StatusCode:  resp.StatusCode,
			LastChecked: currentTime,
		}
	}
}

func checkAllEndpoints() {
	for _, endpoint := range endpoints {
		checkEndpoint(endpoint)
	}
}
