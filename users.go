package main

import ()

type User struct {
	Id       int      `json:"id"`
	Accounts Accounts `json:"accounts"`
}
