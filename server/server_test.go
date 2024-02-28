package server

import (
	"testing"
	"time"
)

func TestClientStatusSetStatus(t *testing.T) {
	clientStatus := ClientStatus{Status: make(map[string]time.Time)}
	clientName := "test"
	dt := time.Now()
	clientStatus.setStatus(clientName, dt)
	if clientStatus.GetStatus(clientName) != dt {
		t.Errorf("Client %s expected timestamp %s, but it was %s", clientName, dt, clientStatus.GetStatus(clientName))
	}
}

func TestClientStatusGetNames(t *testing.T) {
	clientStatus := ClientStatus{Status: make(map[string]time.Time)}
	clientName1 := "test1"
	dt1 := time.Now()
	clientStatus.setStatus(clientName1, dt1)
	clientName2 := "test2"
	dt2 := time.Now()
	clientStatus.setStatus(clientName2, dt2)
	clientNames := clientStatus.GetNames()
	clientNamesFound := 0
	for _, clientName := range clientNames {
		if clientName == clientName1 || clientName == clientName2 {
			clientNamesFound++
		}
	}
	if clientNamesFound != 2 {
		t.Errorf("Looking for client names %s and %s, but found %s", clientName1, clientName2, clientNames)
	}

}
