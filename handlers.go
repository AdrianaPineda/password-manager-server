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

	var account Account

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &account); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	newAccount := CreateAccount(account)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newAccount); err != nil {
		panic(err)
	}
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update account %q\n", html.EscapeString(r.URL.Path))
}

func RemoveAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Remove account %q\n", html.EscapeString(r.URL.Path))
}
