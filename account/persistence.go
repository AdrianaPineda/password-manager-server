package account

import (
	"database/sql"
	"log"
)

type AccountDAO struct {
	err error
}

func CreateAccount() {

}

func GetAccounts(db *sql.DB, userId int) (Accounts, error) {

	rows, err := db.Query("select * from accounts where user_id = ?", userId)
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
		err := rows.Scan(&id, &username, &password, &url)
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
