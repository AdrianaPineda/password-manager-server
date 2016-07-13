package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users/{userId}/accounts", GetAccounts).Methods("GET")
	router.HandleFunc("/users/{userId}/accounts", AddAccount).Methods("POST")
	router.HandleFunc("/users/{userId}/accounts", UpdateAccount).Methods("PUT")
	router.HandleFunc("/users/{userId}/accounts", RemoveAccount).Methods("DELETE")

	fmt.Printf("API initialized\n")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get accounts %q\n", html.EscapeString(r.URL.Path))

	account := Account{Username: "test-username", Password: "test-password", Url: "test-url"}

	json.NewEncoder(w).Encode(account)
}

func AddAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add account %q\n", html.EscapeString(r.URL.Path))
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update account %q\n", html.EscapeString(r.URL.Path))
}

func RemoveAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remove account %q\n", html.EscapeString(r.URL.Path))
}
