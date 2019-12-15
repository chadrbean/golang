// https://www.youtube.com/watch?v=t96hBT53S4U

package main

import (
	"mux"
	"log"
	"net/http"
    "encoding/json"
)

type Person struct {
	ID string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`

}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var person Person
    _ = json.NewDecoder(req.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

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
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Missoula", State: "Montana"}})
	people = append(people, Person{ID:"1", Firstname:"Nic", Lastname:"Raboy"})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("Get")
	router.HandleFunc("/people/{id}", GetPeopleEndpoint).Methods("Get")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("Post")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("Post")
	log.Fatal(http.ListenAndServe(":12345", router))

}