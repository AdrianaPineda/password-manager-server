package account

import (
	"encoding/json"
	"fmt"
	infrast "github.com/AdrianaPineda/password-manager-server/infrastructure"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type AccountAPI struct {
	AccountBusiness AccountBusiness
}

// Constants
const userIdFromUrl = "userId"
const accountIdFromUrl = "accountId"

func getParamAsIntFromRoute(key string, r *http.Request) (int, error) {

	vars := mux.Vars(r)
	paramAsString := vars[key]
	paramAsInt, err := strconv.Atoi(paramAsString)

	return paramAsInt, err
}

func getStringToIntError(value string) infrast.ErrorResponse {

	errorMessage := fmt.Sprintf("Error parsing %s: is not a valid int", value)

	errorResponse := infrast.ErrorResponse{Message: errorMessage}

	return errorResponse
}

// Account related handlers
func (api AccountAPI) GetAccounts(w http.ResponseWriter, r *http.Request) {

	userIdAsInt, err := getParamAsIntFromRoute(userIdFromUrl, r)

	if err != nil {

		errorResponse := getStringToIntError(userIdFromUrl)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)

		return
	}

	accounts, err := api.AccountBusiness.GetAccountsForUser(userIdAsInt)

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

func (api AccountAPI) CreateAccount(w http.ResponseWriter, r *http.Request) {

	userIdAsInt, err := getParamAsIntFromRoute(userIdFromUrl, r)

	if err != nil {

		errorResponse := getStringToIntError(userIdFromUrl)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)

		return
	}

	var currentAccount Account

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

	newAccountId, createError := api.AccountBusiness.CreateAccount(currentAccount, userIdAsInt)
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

func (api AccountAPI) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var currentAccount Account

	vars := mux.Vars(r)
	accountIdAsString := vars[accountIdFromUrl]
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

	updatedAccount, updateError := api.AccountBusiness.UpdateAccount(currentAccount)

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

func (api AccountAPI) DeleteAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	accountIdAsString := vars[accountIdFromUrl]
	accountIdAsInt, err := strconv.Atoi(accountIdAsString)

	if err != nil {
		panic(err)
	}

	if err := api.AccountBusiness.DeleteAccount(accountIdAsInt); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
