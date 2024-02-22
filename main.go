package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"power-monitor/client"
	"power-monitor/server"
	"time"
)

func runServer() {
	clientStatus := server.ClientStatus{Status: make(map[string]time.Time)}
	http.HandleFunc("/", server.Index)
	http.HandleFunc("/checkin", server.ClientCheckIn(&clientStatus))
	listenPort := 8000
	log.Printf("Server is listening on port %d.", listenPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
	log.Printf("%s", err)
}

func runClient(serverHostname string, serverPort int, clientName string, requestDelay int) {
	requestURL := fmt.Sprintf("http://%s:%d/checkin", serverHostname, serverPort)
	log.Printf("Using check-in URL %s", requestURL)
	for {
		// Perform checkin request
		_, err := client.CheckIn(requestURL, clientName)
		if err != nil {
			log.Printf("Error making http request: %s\n", err)
		} else {
			log.Println("Successfully checked in with server.")
		}
		// Sleep for requestDelay seconds
		time.Sleep(time.Duration(requestDelay) * time.Second)
	}
}

func main() {
	// Parse command line arguments
	serverPtr := flag.Bool("server", false, "Run the server.")
	clientPtr := flag.Bool("client", false, "Run the client.")
	// Args only for client
	serverHostnamePtr := flag.String("hostname", "localhost", "Hostname for the server to check in with.")
	serverPortPtr := flag.Int("port", 8000, "Port the server is listening on.")
	clientNamePtr := flag.String("name", "test", "Name for this client. It should be unique.")
	requestDelayPtr := flag.Int("delay", 60, "Delay between checking requests, in seconds.")
	flag.Parse()

	if *serverPtr && *clientPtr {
		log.Println("Cannot simultaneously run the server and the client.")
		os.Exit(1)
	} else if *serverPtr {
		log.Println("Starting server...")
		runServer()
	} else if *clientPtr {
		log.Println("Starting client...")
		runClient(*serverHostnamePtr, *serverPortPtr, *clientNamePtr, *requestDelayPtr)
	} else {
		log.Println("Please specify if the client (-client) or server (-server) should be run.")
	}

}
