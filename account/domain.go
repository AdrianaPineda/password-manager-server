package account

import (
	"database/sql"
)

type AccountBusiness struct {
	AccountDAO AccountDAO
	Database   *sql.DB
}

func (business AccountBusiness) GetAccountsForUser(userId int) (Accounts, error) {
	return business.AccountDAO.GetAccountsOfUserFromDB(business.Database, userId)
}

func (business AccountBusiness) CreateAccount(account Account, userId int) (int, error) {
	return business.AccountDAO.CreateAccountInDB(business.Database, account, userId)
}

func (business AccountBusiness) UpdateAccount(account Account) (Account, error) {
	return business.AccountDAO.UpdateAccountInDB(business.Database, account)
}

func (business AccountBusiness) DeleteAccount(accountId int) error {
	return business.AccountDAO.DeleteAccountFromDB(business.Database, accountId)
}
