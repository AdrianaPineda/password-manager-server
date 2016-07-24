package account

import ()

type Account struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

type Accounts []Account
