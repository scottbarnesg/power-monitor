package main

import (
	"fmt"
	"net/http"
	"power-monitor/server"
	"time"
)

func runServer() {
	clientStatus := server.ClientStatus{Status: make(map[string]time.Time)}
	http.HandleFunc("/", server.Index)
	http.HandleFunc("/checkin", server.ClientCheckIn(&clientStatus))
	listenPort := 8000
	fmt.Printf("Server is listening on port %d.", listenPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
	fmt.Printf("%s", err)
}

func main() {
	runServer()
}
