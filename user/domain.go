package user

import (
	"database/sql"
	"errors"
)

type UserBusiness struct {
	Database *sql.DB
	UserDAO  UserDAO
}

func (userBusiness UserBusiness) CreateUser(user User) (int, error) {
	return userBusiness.UserDAO.CreateUserInDB(userBusiness.Database, user)
}

func (userBusiness UserBusiness) GetUser(userId int) (User, error) {
	return userBusiness.UserDAO.GetUserFromDB(userBusiness.Database, userId)
}

func (userBusiness UserBusiness) GetUsers() ([]*User, error) {
	return userBusiness.UserDAO.GetUsersFromDB(userBusiness.Database)
}

func (userBusiness UserBusiness) UpdateUser(user User) (User, error) {

	_, err := userBusiness.UserDAO.GetUserFromDB(userBusiness.Database, user.Id)
	if err != nil {
		return user, errors.New("User not found")
	}

	return userBusiness.UserDAO.UpdateUserInDB(userBusiness.Database, user)
}

func (userBusiness UserBusiness) DeleteUser(userId int) error {
	return userBusiness.UserDAO.DeleteUserFromDB(userBusiness.Database, userId)
}
