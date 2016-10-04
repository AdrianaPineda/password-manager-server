package user

import (
	database "github.com/AdrianaPineda/password-manager-server/database"
)

func GetUsers() ([]*User, error) {

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

func AddUser(user User) (int, error) {

	var userid int
	err := database.DB.QueryRow("INSERT INTO users(username, password) VALUES($1, $2) RETURNING id", user.UserName, user.Password).Scan(&userid)

	return userid, err
}
