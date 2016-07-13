package main

import ()

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

type Accounts []Account
