package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Constants
const userIdFromUrl = "userId"

type UserAPI struct {
	UserBusiness UserBusiness
}

func (userAPI UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) {

	var currentUser User

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

	newUser, createError := userAPI.UserBusiness.CreateUser(currentUser)

	if createError != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)

		apiResponse := ErrorAPIResponse{Message: "User couldnt be created"}

		log.Printf("Error %v", apiResponse)

		if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
			panic(err)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	apiResponse := SuccessAPIResponse{Message: "User created successfully", User: newUser}
	if err := json.NewEncoder(w).Encode(apiResponse); err != nil {
		panic(err)
	}
}

func (userAPI UserAPI) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	currentUserIdAsString := vars[userIdFromUrl]
	currentUserIdAsInt, err := strconv.Atoi(currentUserIdAsString)

	if err != nil {
		panic(err)
	}

	user, getError := userAPI.UserBusiness.GetUser(currentUserIdAsInt)

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

func (userAPI UserAPI) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, getError := userAPI.UserBusiness.GetUsers()

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

func (userAPI UserAPI) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var currentUser User

	vars := mux.Vars(r)
	currentUserIdAsString := vars[userIdFromUrl]
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

	userUpdated, updateError := userAPI.UserBusiness.UpdateUser(currentUser)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if updateError != nil {

		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(updateError.Error()); err != nil {
			panic(err)
		}

	} else {

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(userUpdated); err != nil {
			panic(err)
		}

	}

}

func (userAPI UserAPI) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	currentUserIdAsString := vars[userIdFromUrl]
	currentUserIdAsInt, err := strconv.Atoi(currentUserIdAsString)

	if err != nil {
		panic(err)
	}

	updateError := userAPI.UserBusiness.DeleteUser(currentUserIdAsInt)

	if updateError != nil {
		panic(updateError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
