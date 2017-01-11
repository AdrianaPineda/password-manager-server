package user

import (
	"database/sql"
	"errors"
)

type UserBusiness struct {
	Database *sql.DB
	UserDAO  UserDAO
}

func (userBusiness UserBusiness) CreateUser(user User) (User, error) {

	if user.Password == "" {
		return User{}, errors.New("Invalid Password")
	}

	if user.UserName == "" {
		return User{}, errors.New("Invalid username")
	}

	userid, err := userBusiness.UserDAO.CreateUserInDB(userBusiness.Database, user)

	if err == nil {
		user.Id = userid
		return user, nil
	}

	return User{}, err
}

func (userBusiness UserBusiness) GetUser(username string) (User, error) {

	if len(username) <= 0 && username != "" {
		return User{}, errors.New("Invalid username")
	}

	return userBusiness.UserDAO.GetUserFromDB(userBusiness.Database, username)
}

func (userBusiness UserBusiness) GetUsers() ([]*User, error) {
	return userBusiness.UserDAO.GetUsersFromDB(userBusiness.Database)
}

func (userBusiness UserBusiness) UpdateUser(user User) (User, error) {

	_, err := userBusiness.UserDAO.GetUserFromDB(userBusiness.Database, user.UserName)
	if err != nil {
		return user, errors.New("User not found")
	}

	return userBusiness.UserDAO.UpdateUserInDB(userBusiness.Database, user)
}

func (userBusiness UserBusiness) DeleteUser(username string) error {

	_, err := userBusiness.UserDAO.GetUserFromDB(userBusiness.Database, username)
	if err != nil {
		return errors.New("User not found")
	}

	return userBusiness.UserDAO.DeleteUserFromDB(userBusiness.Database, username)
}
