package user

import (
	account "github.com/AdrianaPineda/password-manager-server/account"
)

type User struct {
	Id       int              `json:"id"`
	UserName string           `json:"username"`
	Password string           `json:"password"`
	Accounts account.Accounts `json:"accounts"`
}

type Users []User
