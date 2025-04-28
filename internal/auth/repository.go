package auth

import "github.com/bekhuli/go-todo/internal/db"

func CreateUser(username, password string) error {
	_, err := db.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	return err
}

func GetUserByUsername(username string) (*User, error) {
	row := db.DB.QueryRow("SELECT * FROM users WHERE username = $1", username)

	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, err
}
