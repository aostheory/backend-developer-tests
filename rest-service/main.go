package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
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

func returnPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	uuid, err := uuid.FromString(key)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	person, err := models.FindPersonByID((uuid))
	if err != nil {
		log.Printf("%s", err.Error())
		http.Error(w, "404 Not Found", http.StatusInternalServerError)
		return
	}
	fmt.Println("Endpoint Hit: returnPersonById")
	json.NewEncoder(w).Encode(person)
}

func returnPersonByName(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	person := models.FindPeopleByName(firstName, lastName)
	if len(person) == 0 {
		http.Error(w, "404 Not Found", http.StatusInternalServerError)
	}
	fmt.Println("Endpoint Hit: returnPersonByName")
	json.NewEncoder(w).Encode(person)
}

func returnPersonByPhoneNumber(w http.ResponseWriter, r *http.Request) {
	phoneNumber := r.URL.Query().Get("phone_number")

	person := models.FindPeopleByPhoneNumber(phoneNumber)
	if len(person) == 0 {
		http.Error(w, "404 Not Found", http.StatusInternalServerError)
	}
	fmt.Println("Endpoint Hit: returnPersonByPhoneNumber")
	json.NewEncoder(w).Encode(person)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/people/{id}", returnPersonById)
	myRouter.HandleFunc("people", returnPersonByName)
	myRouter.HandleFunc("people", returnPersonByPhoneNumber)
	myRouter.HandleFunc("/people", returnAllPeople)
	myRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	handleRequests()
}
