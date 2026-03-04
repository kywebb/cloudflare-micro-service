package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// AuditEvent represents an enterprise log
type AuditEvent struct {
	User		string		`json:"user"`
	Action		string		`json:"action"`
	Time		time.Time	`json:"timestamp"`
	Success 	bool		`json:"success"`				
}

func accessHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Simulate an Enterprise RBAC check
	userRole := r.Header.Get("X-User-Role") // Mocking a header check

	event := AuditEvent{
		User:		"admin@example.com",
		Action:		"ACCESS_SENSITIVE_DATA",
		Time:		time.Now(),
	}

	if userRole != "admin" {
		event.Success = false
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error: 403 Forbidden - Admin Role Required")
	} else {
		event.Success = true
		fmt.Fprint(w, "Welcome to the Enterprise Dashboard")
	}
	// 2. Log the event (This is where Kafka would usually come in)
	logData, _ := json.Marshal(event)
	fmt.Printf("AUDIT LOG SENT: %s\n", string(logData))
}

func main() {
	http.HandleFunc("/api/v1/access", accessHandler)

	fmt.Println("Enterprise Readiness Service starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}