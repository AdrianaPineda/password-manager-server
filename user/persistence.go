package user

import (
	"database/sql"
	"log"
)

type UserDAO struct {
}

// CREATE
func (userDAO UserDAO) CreateUserInDB(database *sql.DB, user User) (int, error) {

	var userid int
	err := database.QueryRow("INSERT INTO users(username, password) VALUES($1, $2) RETURNING id", user.UserName, user.Password).Scan(&userid)

	return userid, err
}

// READ
func (userDAO UserDAO) GetUserFromDB(database *sql.DB, userId int) (User, error) {

	var user User
	err := database.QueryRow("SELECT * FROM users WHERE id = $1", userId).Scan(&user.Id, &user.UserName, &user.Password)

	return user, err
}

func (userDAO UserDAO) GetUsersFromDB(database *sql.DB) ([]*User, error) {

	rows, err := database.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Id, &user.UserName, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// UPDATE
func (userDAO UserDAO) UpdateUserInDB(database *sql.DB, user User) (User, error) {

	smt, err := database.Prepare("UPDATE users SET username = $1, password = $2 WHERE id = $3")

	if err != nil {
		log.Fatal(err)
	}

	defer smt.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = smt.Exec(user.UserName, user.Password, user.Id)

	return user, err

}

// DELETE
func (userDAO UserDAO) DeleteUserFromDB(database *sql.DB, userId int) error {

	smt, err := database.Prepare("DELETE FROM users WHERE id = $1")

	if err != nil {
		log.Fatal(err)
	}

	defer smt.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = smt.Exec(userId)

	return err
}
