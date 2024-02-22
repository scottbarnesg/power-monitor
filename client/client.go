package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func CheckIn(requestURL string, clientName string) (*http.Response, error) {
	// Format payload
	payload := make(map[string]string)
	payload["name"] = clientName
	jsonPayload, _ := json.Marshal(payload)
	// Send request
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(jsonPayload))
	return resp, err
}
