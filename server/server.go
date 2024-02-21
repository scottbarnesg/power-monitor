package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ClientPayload struct {
	Name string `json:"name"`
}

type Response struct {
	Status string `json:"status"`
}

// TODO: Make this thread-safe
type ClientStatus struct {
	Status map[string]time.Time
}

func Index(w http.ResponseWriter, r *http.Request) {
	// Validate request type
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Generate response
	w.Header().Set("Content-Type", "application/json")
	payload := make(map[string]string)
	payload["status"] = "OK"
	json.NewEncoder(w).Encode(payload)
}

func ClientCheckIn(clientStatus *ClientStatus) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate request type
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			// TODO: Update this to return a JSON payload
			return
		}
		// Get client payload from body
		clientPayload := &ClientPayload{}
		err := json.NewDecoder(r.Body).Decode(clientPayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			// TODO: Update this to return a JSON payload
			return
		}
		// TODO: Get client name and the current timestamp and add it to a shared data structure somewhere.
		dt := time.Now()
		fmt.Printf("Client %s checked in at %s.\n", clientPayload.Name, dt.String())
		clientStatus.Status[clientPayload.Name] = dt
		fmt.Printf("%s\n", clientStatus.Status)
		// Generate response
		w.Header().Set("Content-Type", "application/json")
		payload := make(map[string]string)
		payload["status"] = "OK"
		json.NewEncoder(w).Encode(payload)
	}
}
