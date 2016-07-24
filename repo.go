package main

import (
	"database/sql"
	"fmt"
	account "github.com/AdrianaPineda/password-manager-server/account"
	_ "github.com/go-sql-driver/mysql"
)

var driverName = "mysql"
var dbCredentials = "user:password@tcp(127.0.0.1:3306)/hello"

var currentAccountId int

var accounts account.Accounts

func FindAccountById(id int) account.Account {
	for _, t := range accounts {
		if t.Id == id {
			return t
		}
	}

	return account.Account{}
}

func CreateAccount(account account.Account) account.Account {

	db, err := sql.Open(driverName, dbCredentials)

	if err != nil {
		// log.Fatal(err)
	}

	defer db.Close()

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

func GetAllAccounts() account.Accounts {
	return accounts
}

func UpdateSingleAccount(account account.Account) account.Account {

	// if err := DestroyAccount(account.Id); err != nil {
	// 	return account.Account{}
	// }

	accounts = append(accounts, account)

	return account
}
