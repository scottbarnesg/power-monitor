package monitor

import (
	"fmt"
	"log"
	"power-monitor/config"
	"power-monitor/email"
	"power-monitor/server"
	"time"
)

type AlertStatus struct {
	ClientName string
	IsOnline   bool
}

func MonitorClientStatus(clientStatus *server.ClientStatus, config *config.Config, alertThreshold time.Duration) {
	alertStatuses := make(map[string]AlertStatus)
	log.Printf("Monitoring client statuses with Alert Threshold %s.\n", alertThreshold.String())
	for {
		for _, clientName := range clientStatus.GetNames() {
			// Create this client's AlertStatus if it doesn't exist
			if _, exists := alertStatuses[clientName]; !exists {
				alertStatuses[clientName] = AlertStatus{ClientName: clientName, IsOnline: true}
				// This is a new client. Send an email indicating we detected a new client.
				log.Printf("Identified new client: %s.\n", clientName)
				emailSubject := fmt.Sprintf("Power Monitor Alert: New Client - %s.", clientName)
				emailBody := fmt.Sprintf("Power Monitor detected new client: %s.\n", clientName)
				email.SendEmail(config.From, config.To, emailSubject, emailBody, config.From, config.Password)
			}
			// Calculate how long its been since the client last checked in.
			clientLastCheckin := clientStatus.GetStatus(clientName)
			timeSinceLastCheckin := time.Since(clientLastCheckin)
			// Handle case where the client's last checkin is older than alertThreshold
			if timeSinceLastCheckin > alertThreshold {
				// If we havent sent an alert, send one and update this client's AlertStatus
				clientAlertStatus := alertStatuses[clientName]
				if clientAlertStatus.IsOnline {
					// Update this client's AlertStatus
					clientAlertStatus.IsOnline = false
					// Send an "alert" email
					log.Printf("Client %s is offline. Last checkin was %s.\n", clientName, clientLastCheckin.String())
					emailSubject := fmt.Sprintf("Power Monitor Alert: %s is Offline", clientName)
					emailBody := fmt.Sprintf("Power Monitor detected that %s is offline. It last checked in at %s\n.", clientName, clientLastCheckin.String())
					email.SendEmail(config.From, config.To, emailSubject, emailBody, config.From, config.Password)
					// Update AlertStatuses for this client
					alertStatuses[clientName] = clientAlertStatus
				}
			} else { // Handle case where client's last checkin is more recent than AlertThreshold
				// If this is marked as offline, mark it as back online and send a "resolved" email.
				clientAlertStatus := alertStatuses[clientName]
				// If this client is marked as offline, it has just come back online.
				if !clientAlertStatus.IsOnline {
					// Send an "issue resolved" email.
					log.Printf("Client %s is back online.\n", clientName)
					emailSubject := fmt.Sprintf("Power Monitor Alert: %s is back Online", clientName)
					emailBody := fmt.Sprintf("Power Monitor detected that %s is back online. It last checked in at %s\n.", clientName, clientLastCheckin.String())
					email.SendEmail(config.From, config.To, emailSubject, emailBody, config.From, config.Password)
					// Update AlertStatuses for this client
					clientAlertStatus.IsOnline = true
					alertStatuses[clientName] = clientAlertStatus
				} else {
					log.Printf("Client %s is healthy. Last checkin was %s.\n", clientName, clientLastCheckin)
				}
			}
		}
		time.Sleep(alertThreshold)
	}
}
