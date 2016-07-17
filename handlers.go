package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get accounts %q\n", html.EscapeString(r.URL.Path))

	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)

	accounts = GetAllAccounts()

	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		panic(err)
	}
}

func AddAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add account %q\n", html.EscapeString(r.URL.Path))

	vars := mux.Vars(r)
	userId := vars["user_id"]
	fmt.Fprintln(w, "Add account for user: ", userId)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update account %q\n", html.EscapeString(r.URL.Path))
}

func RemoveAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remove account %q\n", html.EscapeString(r.URL.Path))
}
