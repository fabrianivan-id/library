package models

import "errors"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func GetUserByUsername(username string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, password, role FROM users WHERE username = ?", username).
		Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	if err != nil {
		return User{}, errors.New("user not found")
	}
	return user, nil
}
