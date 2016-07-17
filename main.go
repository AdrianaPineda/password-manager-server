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

	account_1 := Account{Username: "test-username-1", Password: "test-password-1", Url: "test-url-1"}
	account_2 := Account{Username: "test-username-2", Password: "test-password-2", Url: "test-url-2"}
	accounts := Accounts{account_1, account_2}

	json.NewEncoder(w).Encode(accounts)
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
