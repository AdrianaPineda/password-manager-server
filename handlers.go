package main

import (
	"encoding/json"
	account "github.com/AdrianaPineda/password-manager-server/account"
	user "github.com/AdrianaPineda/password-manager-server/user"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Account related handlers
func GetAccounts(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userIdAsString := vars["userId"]
	userIdAsInt, err := strconv.Atoi(userIdAsString)

	if err != nil {
		panic(err)
	}

	accounts, err := account.GetAccountsOfUserFromDB(userIdAsInt)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		panic(err)
	}
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userIdAsString := vars["userId"]
	userIdAsInt, err := strconv.Atoi(userIdAsString)

	if err != nil {
		panic(err)
	}

	var currentAccount account.Account

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &currentAccount); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	newAccountId, createError := account.CreateAccountInDB(currentAccount, userIdAsInt)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if createError != nil {

		log.Printf("Error %v", createError)
		w.WriteHeader(http.StatusBadRequest)

		if err := json.NewEncoder(w).Encode(createError); err != nil {
			panic(err)
		}

	} else {
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(newAccountId); err != nil {
			panic(err)
		}
	}
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var currentAccount account.Account

	vars := mux.Vars(r)
	accountIdAsString := vars["accountId"]
	accountIdAsInt, err := strconv.Atoi(accountIdAsString)

	if err != nil {
		panic(err)
	}

	currentAccount.Id = accountIdAsInt

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &currentAccount); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	updatedAccount, updateError := account.UpdateAccountInDB(currentAccount)

	if updateError != nil {
		log.Printf("%v", updateError)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err := json.NewEncoder(w).Encode(updatedAccount); err != nil {
		panic(err)
	}
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	accountIdAsString := vars["accountId"]
	accountIdAsInt, err := strconv.Atoi(accountIdAsString)

	if err != nil {
		panic(err)
	}

	if err := account.DeleteAccountFromDB(accountIdAsInt); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// USER related handlers
func CreateUser(w http.ResponseWriter, r *http.Request) {

	var currentUser user.User

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &currentUser); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	newUserId, createError := user.CreateUserInDB(currentUser)

	if createError != nil {
		panic(createError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newUserId); err != nil {
		panic(err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	currentUserIdAsString := vars["userId"]
	currentUserIdAsInt, err := strconv.Atoi(currentUserIdAsString)

	if err != nil {
		panic(err)
	}

	user, getError := user.GetUserFromDB(currentUserIdAsInt)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if getError != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(user); err != nil {
			panic(err)
		}
	}

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	users, getError := user.GetUsersFromDB()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if getError != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(users); err != nil {
			panic(err)
		}
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var currentUser user.User

	vars := mux.Vars(r)
	currentUserIdAsString := vars["userId"]
	currentUserIdAsInt, err := strconv.Atoi(currentUserIdAsString)

	currentUser.Id = currentUserIdAsInt

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &currentUser); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	userUpdated, updateError := user.UpdateUserInDB(currentUser)

	if updateError != nil {
		panic(updateError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(userUpdated); err != nil {
		panic(err)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var currentUser user.User

	vars := mux.Vars(r)
	currentUserIdAsString := vars["userId"]
	currentUserIdAsInt, err := strconv.Atoi(currentUserIdAsString)

	currentUser.Id = currentUserIdAsInt

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &currentUser); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	updateError := user.DeleteUserFromDB(currentUser)

	if updateError != nil {
		panic(updateError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
