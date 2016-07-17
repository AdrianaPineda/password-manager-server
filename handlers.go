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

	account_1 := Account{Username: "test-username-1", Password: "test-password-1", Url: "test-url-1"}
	account_2 := Account{Username: "test-username-2", Password: "test-password-2", Url: "test-url-2"}
	accounts := Accounts{account_1, account_2}

	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)

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
