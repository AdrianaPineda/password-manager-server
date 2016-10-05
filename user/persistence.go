package user

import (
	database "github.com/AdrianaPineda/password-manager-server/database"
	"log"
)

// CREATE
func CreateUserInDB(user User) (int, error) {

	var userid int
	err := database.DB.QueryRow("INSERT INTO users(username, password) VALUES($1, $2) RETURNING id", user.UserName, user.Password).Scan(&userid)

	return userid, err
}

// READ
func GetUserFromDB(userId int) (User, error) {

	var user User
	err := database.DB.QueryRow("SELECT * FROM users WHERE id = $1", userId).Scan(&user.Id, &user.UserName, &user.Password)

	return user, err
}

func GetUsersFromDB() ([]*User, error) {

	rows, err := database.DB.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*User, 0)

	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Id, &user.UserName)
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
func UpdateUserInDB(user User) (User, error) {

	smt, err := database.DB.Prepare("UPDATE users SET username = $1, password = $2 WHERE id = $3")

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
func DeleteUserFromDB(user User) error {

	smt, err := database.DB.Prepare("DELETE FROM users WHERE id = $1")

	if err != nil {
		log.Fatal(err)
	}

	defer smt.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, err = smt.Exec(user.Id)

	return err
}
