package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Person is just structure...
type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

//Address of the person
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

//GetPersonEndpoint just try to get a single record
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

//GetPeopleEndpoint show all person
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

//CreatePersonEndpoint just create
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)

}

//DeletePersonEndpoint just delete
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)

}
func main() {
	fmt.Println("Server Started.......")
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "Mohammad", LastName: "Asiuzzaman", Address: &Address{City: "Rajshahi", State: "Bangladesh"}})
	people = append(people, Person{ID: "2", FirstName: "Farhat", LastName: "Zim", Address: &Address{City: "Khustia", State: "Bangladesh"}})
	people = append(people, Person{ID: "3", FirstName: "kjjjjsdjhsh", LastName: "asdfhgjgj", Address: &Address{City: "blabla", State: "Bangladesh"}})

	fmt.Println(len(people))
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
