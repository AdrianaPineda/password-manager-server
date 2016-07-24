package main

import (
	account "github.com/AdrianaPineda/password-manager-server/account"
)

type User struct {
	Id       int              `json:"id"`
	Accounts account.Accounts `json:"accounts"`
}
