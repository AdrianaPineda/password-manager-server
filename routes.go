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
	Route{Name: "Get accounts", Method: "GET", Pattern: "/users/{userId}/accounts", HandlerFunc: GetAccounts},
	Route{Name: "Add Account", Method: "POST", Pattern: "/users/{userId}/accounts", HandlerFunc: AddAccount},
	Route{Name: "Update Account", Method: "PUT", Pattern: "/users/{userId}/accounts", HandlerFunc: UpdateAccount},
	Route{Name: "Remove Account", Method: "DELETE", Pattern: "/users/{userId}/accounts/{accountId:[0-9]+}", HandlerFunc: RemoveAccount},
}
