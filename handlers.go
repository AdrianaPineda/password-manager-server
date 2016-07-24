package main

import (
	"encoding/json"
	account "github.com/AdrianaPineda/password-manager-server/account"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	accounts = GetAllAccounts()

	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		panic(err)
	}
}

func AddAccount(w http.ResponseWriter, r *http.Request) {

	var account account.Account

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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newAccount); err != nil {
		panic(err)
	}
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var account account.Account

	vars := mux.Vars(r)
	accountIdAsString := vars["accountId"]
	accountIdAsInt, err := strconv.Atoi(accountIdAsString)

	if err != nil {
		panic(err)
	}

	account.Id = accountIdAsInt

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

	updatedAccount := UpdateSingleAccount(account)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(updatedAccount); err != nil {
		panic(err)
	}
}

func RemoveAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	accountIdAsString := vars["accountId"]
	accountIdAsInt, err := strconv.Atoi(accountIdAsString)

	if err != nil {
		panic(err)
	}

	if err := DestroyAccount(accountIdAsInt); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
