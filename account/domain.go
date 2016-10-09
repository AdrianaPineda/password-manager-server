package account

import (
	"database/sql"
	"errors"
	user "github.com/AdrianaPineda/password-manager-server/user"
)

type AccountBusiness struct {
	AccountDAO   AccountDAO
	Database     *sql.DB
	UserBusiness user.UserBusiness
}

func (business AccountBusiness) GetAccountsForUser(userId int) (Accounts, error) {
	return business.AccountDAO.GetAccountsOfUserFromDB(business.Database, userId)
}

func (business AccountBusiness) CreateAccount(account Account, userId int) (int, error) {

	// validate user exists
	if business.doesUserExists(userId) {
		return business.AccountDAO.CreateAccountInDB(business.Database, account, userId)
	}

	return -1, errors.New("User not found")

}

func (business AccountBusiness) UpdateAccount(account Account) (Account, error) {
	return business.AccountDAO.UpdateAccountInDB(business.Database, account)
}

func (business AccountBusiness) DeleteAccount(accountId int) error {
	return business.AccountDAO.DeleteAccountFromDB(business.Database, accountId)
}

func (business AccountBusiness) doesUserExists(userId int) bool {

	_, err := business.UserBusiness.GetUser(userId)
	if err == nil {
		return true
	}

	return false
}
