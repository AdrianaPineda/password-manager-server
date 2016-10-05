package account

import (
	database "github.com/AdrianaPineda/password-manager-server/database"
	"log"
)

type AccountDAO struct {
	err error
}

// CREATE
func CreateAccountInDB(account Account, userId int) (int, error) {

	var accountId int
	err := database.DB.QueryRow("INSERT INTO accounts(username, password, url, userId) VALUES($1, $2, $3, $4) RETURNING id", account.Username, account.Password, account.Url, userId).Scan(&accountId)

	return accountId, err

}

// READ
func GetAccountsOfUserFromDB(userId int) (Accounts, error) {

	rows, err := database.DB.Query("SELECT * from accounts where userId = $1", userId)
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
func UpdateAccountInDB(account Account) (Account, error) {

	smt, err := database.DB.Prepare("UPDATE accounts SET username = $1, password = $2, url = $3 WHERE id = $4")

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
func DeleteAccountFromDB(accountId int) error {

	smt, err := database.DB.Prepare("DELETE FROM accounts WHERE id = $1")

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
