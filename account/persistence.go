package account

import (
	"database/sql"
	"log"
)

type AccountDAO struct {
}

// CREATE
func (dao AccountDAO) CreateAccountInDB(database *sql.DB, account Account, userId int) (int, error) {

	var accountId int
	err := database.QueryRow("INSERT INTO accounts(username, password, url, userId) VALUES($1, $2, $3, $4) RETURNING id", account.Username, account.Password, account.Url, userId).Scan(&accountId)

	return accountId, err

}

// READ
func (dao AccountDAO) GetAccountsOfUserFromDB(database *sql.DB, userId int) (Accounts, error) {

	rows, err := database.Query("SELECT * from accounts where userId = $1", userId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var currentAccounts Accounts

	var (
		id       int
		username string
		password string
		url      string
	)
	for rows.Next() {
		err := rows.Scan(&id, &username, &password, &url, &userId)
		if err != nil {
			log.Fatal(err)
		}
		currentAccount := Account{Id: id, Username: username, Password: password, Url: url}
		currentAccounts = append(currentAccounts, currentAccount)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return currentAccounts, err
}

// UPDATE
func (dao AccountDAO) UpdateAccountInDB(database *sql.DB, account Account) (Account, error) {

	smt, err := database.Prepare("UPDATE accounts SET username = $1, password = $2, url = $3 WHERE id = $4")

	if err != nil {
		log.Fatal(err)
	}

	defer smt.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = smt.Exec(account.Username, account.Password, account.Url, account.Id)

	return account, err

}

// DELETE
func (dao AccountDAO) DeleteAccountFromDB(database *sql.DB, accountId int) error {

	smt, err := database.Prepare("DELETE FROM accounts WHERE id = $1")

	if err != nil {
		log.Fatal(err)
	}

	defer smt.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = smt.Exec(accountId)

	return err
}
