package main

import (
	account "github.com/AdrianaPineda/password-manager-server/account"
	user "github.com/AdrianaPineda/password-manager-server/user"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func getUserRoutes(userAPI user.UserAPI) Routes {

	userRoutes := Routes{
		Route{Name: "Create user", Method: "POST", Pattern: "/users", HandlerFunc: userAPI.CreateUser},
		Route{Name: "Get user", Method: "GET", Pattern: "/users/{userId:[0-9]+}", HandlerFunc: userAPI.GetUser},
		Route{Name: "Get users", Method: "GET", Pattern: "/users", HandlerFunc: userAPI.GetUsers},
		Route{Name: "Update user", Method: "PUT", Pattern: "/users/{userId:[0-9]+}", HandlerFunc: userAPI.UpdateUser},
		Route{Name: "Delete user", Method: "DELETE", Pattern: "/users/{userId:[0-9]+}", HandlerFunc: userAPI.DeleteUser},
	}

	return userRoutes
}

func getAccountRoutes(accountAPI account.AccountAPI) Routes {

	accountRoutes := Routes{

		Route{Name: "Get accounts", Method: "GET", Pattern: "/users/{userId:[0-9]+}/accounts", HandlerFunc: accountAPI.GetAccounts},
		Route{Name: "Add Account", Method: "POST", Pattern: "/users/{userId:[0-9]+}/accounts", HandlerFunc: accountAPI.CreateAccount},
		Route{Name: "Update Account", Method: "PUT", Pattern: "/users/{userId:[0-9]+}/accounts/{accountId:[0-9]+}", HandlerFunc: accountAPI.UpdateAccount},
		Route{Name: "Delete Account", Method: "DELETE", Pattern: "/users/{userId:[0-9]+}/accounts/{accountId:[0-9]+}", HandlerFunc: accountAPI.DeleteAccount},
	}

	return accountRoutes
}
