package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the StackPath Web Service Homepage.")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPeople")
	json.NewEncoder(w).Encode(models.AllPeople())
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/people", returnAllPeople)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	// TODO: Add RESTful web service here
	handleRequests()
}
