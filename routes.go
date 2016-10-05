package main

import (
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
	Route{Name: "Create user", Method: "POST", Pattern: "/users", HandlerFunc: CreateUser},
	Route{Name: "Get user", Method: "GET", Pattern: "/users/{userId:[0-9]+}", HandlerFunc: GetUser},
	Route{Name: "Update user", Method: "PUT", Pattern: "/users/{userId:[0-9]+}", HandlerFunc: UpdateUser},

	Route{Name: "Get accounts", Method: "GET", Pattern: "/users/{userId}/accounts", HandlerFunc: GetAccounts},
	Route{Name: "Add Account", Method: "POST", Pattern: "/users/{userId}/accounts", HandlerFunc: CreateAccount},
	Route{Name: "Update Account", Method: "PUT", Pattern: "/users/{userId}/accounts/{accountId:[0-9]+}", HandlerFunc: UpdateAccount},
	Route{Name: "Remove Account", Method: "DELETE", Pattern: "/users/{userId}/accounts/{accountId:[0-9]+}", HandlerFunc: RemoveAccount},
}
