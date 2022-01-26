package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Creating structure 'dog' with below parameters.
type dog struct {
	Name        string `json:"Name"`
	Color       string `json:"Color"`
	Size        string `json:"Size"`
	Disposition string `json:"Disposition"`
}

// creating allDogs and a dogs structure element with our first dog.
type allDogs []dog

var dogs = allDogs{
	{
		Name:        "Waf",
		Color:       "Black and Orange",
		Size:        "Large",
		Disposition: "Nice but bites people occasionally, has 2 toys and one dog bed, sensitive but fluffy",
	},
	{
		Name:        "Arthur",
		Color:       "Spotty White",
		Size:        "Medium",
		Disposition: "Good heart, Gentle Soul, doesnt bite anyone",
	},
	{
		Name:        "Dasha",
		Color:       "Brown",
		Size:        "Medium",
		Disposition: "Nice, Hungry, Old :( ",
	},
	{
		Name:        "Yozhik",
		Color:       "Black",
		Size:        "Extra Large",
		Disposition: "Really scared of stuff, barky, loves toys and has many toys, bushy tail",
	},
	{
		Name:        "Matisse",
		Color:       "White with Spots",
		Size:        "Extra Super Large",
		Disposition: "Nice, Cat best friend, 6 toes 4 paws, guards stuff (Garden)",
	},
	{
		Name:        "Kimberly",
		Color:       "White",
		Size:        "Medium",
		Disposition: "Shy and Happy",
	},
}

// This function created a a new dog slice in the dog struct and appends it to the existing dog struct.
func createDog(w http.ResponseWriter, r *http.Request) {
	var newDog dog
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter new dog information please!!! Be honest, the dogs reputation is in your hands!")
	}

	json.Unmarshal(reqBody, &newDog)
	dogs = append(dogs, newDog)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newDog)
	fmt.Print(newDog.Name)
}

// gets and returns single dog ingo by name.
func getDogInfo(w http.ResponseWriter, r *http.Request) {
	dogName := mux.Vars(r)["name"]

	for _, singleDog := range dogs {
		if singleDog.Name == dogName {
			json.NewEncoder(w).Encode(singleDog)
			fmt.Print(singleDog.Name)
		}

	}
}

// Finds a dog entry in the struct, by name, and updates it per user input.
func updateDog(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["name"]
	var updatedDog dog

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter new dog description to update dog")
	}
	json.Unmarshal(reqBody, &updatedDog)

	for i, singleDog := range dogs {
		if singleDog.Name == eventID {
			singleDog.Color = updatedDog.Color
			singleDog.Size = updatedDog.Size
			singleDog.Disposition = updatedDog.Disposition
			dogs = append(dogs[:i], singleDog)
			json.NewEncoder(w).Encode(singleDog)
		}
	}
}

// deletes a dog slice from struct, by name.
func deleteDog(w http.ResponseWriter, r *http.Request) {
	dogName := mux.Vars(r)["name"]

	for i, singleDog := range dogs {
		if singleDog.Name == dogName {
			dogs = append(dogs[:i], dogs[i+1:]...)
			fmt.Fprintf(w, "The dog named %v has been deleted successfully", dogName)
		}
	}
}

// returns the entire dogs struct.
func getAllDogs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dogs)

	fmt.Print("all dogs requested")
}

// passes home page info to http request.
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to dog home! Bark! Bark! Bark!  jk they are friendly")
}

// main function with all the dog operations
// initiates mux rouer
func main() {
	// initDogs()
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/dog", createDog).Methods("POST")
	router.HandleFunc("/dogs", getAllDogs).Methods("GET")
	router.HandleFunc("/dogs/{name}", getDogInfo).Methods("GET")
	router.HandleFunc("/dogs/{name}", updateDog).Methods("PATCH")
	router.HandleFunc("/dogs/{name}", deleteDog).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

	//	log.Fatal(http.ListenAndServeTLS(":8080", "certificate.crt", "certificate.key", router))
	//	http.Handle("/", router)
}
