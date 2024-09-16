package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendSlackNotification(endpoint, message string) {
	webhookURL := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	payload := map[string]string{
		"text": fmt.Sprintf("Endpoint %s is down: %s", endpoint, message),
	}

	jsonPayload, _ := json.Marshal(payload)

	_, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Error sending Slack message: %v\n", err)
	}
}
