package main

import (
	"fmt"
)

var currentAccountId int

var accounts Accounts

func FindAccountById(id int) Account {
	for _, t := range accounts {
		if t.Id == id {
			return t
		}
	}

	return Account{}
}

func CreateAccount(account Account) Account {
	currentAccountId += 1
	account.Id = currentAccountId
	accounts = append(accounts, account)
	return account
}

func DestroyAccount(id int) error {

	for i, t := range accounts {
		if t.Id == id {
			accounts = append(accounts[:i], accounts[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find account for with id %d to delete", id)
}

func GetAllAccounts() Accounts {
	return accounts
}
