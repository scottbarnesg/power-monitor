package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type ClientPayload struct {
	Name string `json:"name"`
}

type Response struct {
	Status string `json:"status"`
}

type ClientStatus struct {
	Status map[string]time.Time
	lock   sync.RWMutex
}

func (cs *ClientStatus) setStatus(name string, dt time.Time) {
	cs.lock.RLock()
	defer cs.lock.RUnlock()
	cs.Status[name] = dt
}

func (cs *ClientStatus) getNames() []string {
	keys := make([]string, 0, len(cs.Status))
	cs.lock.RLock()
	defer cs.lock.RUnlock()
	for k := range cs.Status {
		keys = append(keys, k)
	}
	return keys
}

func (cs *ClientStatus) getStatus(name string) time.Time {
	cs.lock.RLock()
	defer cs.lock.RUnlock()
	return cs.Status[name]
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
		log.Printf("Client %s checked in at %s.\n", clientPayload.Name, dt.String())
		clientStatus.setStatus(clientPayload.Name, dt)
		log.Printf("%s\n", clientStatus.Status)
		// Generate response
		w.Header().Set("Content-Type", "application/json")
		payload := make(map[string]string)
		payload["status"] = "OK"
		json.NewEncoder(w).Encode(payload)
	}
}
