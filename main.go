package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Event represents an event
type Event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

// Events global events array
var Events []Event

// landing page
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

// POST: /event - Emits back the created event
func createEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: createEvent")

	reqBody, _ := ioutil.ReadAll(r.Body)
	// unmarshal request body into new Event struct, and append to Events array
	var event Event
	json.Unmarshal(reqBody, &event)
	Events = append(Events, event)

	json.NewEncoder(w).Encode(event)
}

// DELETE: /event/{id} - delete specific event
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: deleteEvent")

	vars := mux.Vars(r)
	id := vars["id"]

	for index, event := range Events {
		if event.ID == id {
			Events = append(Events[:index], Events[index+1:]...)
		}
	}
}

// PUT: /event/{id} - update specific event
func updateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: updateEvent")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var newEvent Event
	json.Unmarshal(reqBody, &newEvent)

	for index, event := range Events {
		if event.ID == newEvent.ID {
			var tempArr = append(Events[:index], Events[index+1:]...)
			Events = append(tempArr, newEvent)
		}
	}
}

// GET: /events - get all events
func getEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getEvents")
	json.NewEncoder(w).Encode(Events)
}

// GET: /event/{id} - get specific event by event ID
func getEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: getEvent")

	vars := mux.Vars(r)
	key := vars["id"]

	for _, event := range Events {
		if event.ID == key {
			json.NewEncoder(w).Encode(event)
		}
	}
}

func handleRequests() {
	// create new instance of mux router
	router := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with router.HandleFunc
	router.HandleFunc("/", homeLink)

	router.HandleFunc("/event", createEvent).Methods("POST") // ordering is important - must be defined before other `/event` endpoint
	router.HandleFunc("/event", updateEvent).Methods("PUT")
	router.HandleFunc("/event/{id}", deleteEvent).Methods("DELETE")
	router.HandleFunc("/event/{id}", getEvent)

	router.HandleFunc("/events", getEvents)

	log.Print("...listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	Events = []Event{
		{ID: "1", Title: "Ballgame", Description: "Virtual baseball"},
		{ID: "2", Title: "Racing", Description: "Virtual IndyCar"},
	}

	handleRequests()
}
