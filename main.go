package main

import (
	"database/sql"
	"errors"
	"fmt"
	account "github.com/AdrianaPineda/password-manager-server/account"
	config "github.com/AdrianaPineda/password-manager-server/config"
	user "github.com/AdrianaPineda/password-manager-server/user"
	"log"
	"net/http"
)

var accountAPI account.AccountAPI
var userAPI user.UserAPI

func main() {

	database, err := config.InitDB("user=adrianaPineda dbname = adrianaPineda sslmode=disable")

	if err != nil {
		panic(err)
	}

	accountAPI := createAccountAPI(database)
	userAPI := createUserAPI(database)

	accountRoutes := getAccountRoutes(accountAPI)
	userRoutes := getUserRoutes(userAPI)

	routes := append(accountRoutes, userRoutes...)

	router := NewRouter(routes)

	fmt.Printf("API initialized\n")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func createAccountAPI(database *sql.DB) (accountAPI account.AccountAPI) {

	accountDAO := account.AccountDAO{}
	accountBusiness := account.AccountBusiness{AccountDAO: accountDAO, Database: database}
	accountAPI = account.AccountAPI{AccountBusiness: accountBusiness}

	if accountAPI == (account.AccountAPI{}) {
		panic(errors.New("API couldnt not be initialized"))
	}

	return accountAPI
}

func createUserAPI(database *sql.DB) (userAPI user.UserAPI) {
	userDAO := user.UserDAO{}
	userBusiness := user.UserBusiness{UserDAO: userDAO, Database: database}
	userAPI = user.UserAPI{UserBusiness: userBusiness}

	if userAPI == (user.UserAPI{}) {
		panic(errors.New("API couldnt not be initialized"))
	}

	return userAPI
}
