package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users/{userId}/accounts", GetAccounts).Methods("GET")
	router.HandleFunc("/accounts", AddAccount).Methods("POST")
	router.HandleFunc("/accounts", UpdateAccount).Methods("PUT")
	router.HandleFunc("/accounts", RemoveAccount).Methods("DELETE")

	fmt.Printf("Hello world\n")

}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get accounts %q", html.EscapeString(r.URL.Path))
}

func AddAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add account %q", html.EscapeString(r.URL.Path))
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update account %q", html.EscapeString(r.URL.Path))
}

func RemoveAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remove account %q", html.EscapeString(r.URL.Path))
}
