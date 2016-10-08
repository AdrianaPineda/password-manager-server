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

var routes = Routes{
	Route{Name: "Create user", Method: "POST", Pattern: "/users", HandlerFunc: user.CreateUser},
	Route{Name: "Get user", Method: "GET", Pattern: "/users/{userId:[0-9]+}", HandlerFunc: user.GetUser},
	Route{Name: "Get users", Method: "GET", Pattern: "/users", HandlerFunc: user.GetUsers},
	Route{Name: "Update user", Method: "PUT", Pattern: "/users/{userId:[0-9]+}", HandlerFunc: user.UpdateUser},
	Route{Name: "Delete user", Method: "DELETE", Pattern: "/users/{userId:[0-9]+}", HandlerFunc: user.DeleteUser},

	Route{Name: "Get accounts", Method: "GET", Pattern: "/users/{userId:[0-9]+}/accounts", HandlerFunc: account.GetAccounts},
	Route{Name: "Add Account", Method: "POST", Pattern: "/users/{userId:[0-9]+}/accounts", HandlerFunc: account.CreateAccount},
	Route{Name: "Update Account", Method: "PUT", Pattern: "/users/{userId:[0-9]+}/accounts/{accountId:[0-9]+}", HandlerFunc: account.UpdateAccount},
	Route{Name: "Delete Account", Method: "DELETE", Pattern: "/users/{userId:[0-9]+}/accounts/{accountId:[0-9]+}", HandlerFunc: account.DeleteAccount},
}
