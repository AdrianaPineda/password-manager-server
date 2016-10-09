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

	if user.Password == "" {
		return -1, errors.New("Invalid Password")
	}

	if user.UserName == "" {
		return -1, errors.New("Invalid username")
	}

	return userBusiness.UserDAO.CreateUserInDB(userBusiness.Database, user)
}

func (userBusiness UserBusiness) GetUser(userId int) (User, error) {

	if userId <= 0 {
		return User{}, errors.New("Invalid id")
	}

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

	_, err := userBusiness.UserDAO.GetUserFromDB(userBusiness.Database, userId)
	if err != nil {
		return errors.New("User not found")
	}

	return userBusiness.UserDAO.DeleteUserFromDB(userBusiness.Database, userId)
}
