package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Event represents an event
type Event struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

// Events global events array
var Events []Event

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getEvents")
	json.NewEncoder(w).Encode(Events)
}

func main() {
	Events = []Event{
		Event{ID: 1, Title: "Ballgame", Description: "Virtual baseball"},
		Event{ID: 2, Title: "Racing", Description: "Virtual IndyCar"},
	}

	http.HandleFunc("/", homeLink)
	http.HandleFunc("/events", getEvents)

	log.Print("...listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
