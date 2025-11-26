package repos

import (
	"myapp-test-btech/cmd/models"
	"myapp-test-btech/configs"
)

func InsertUser(user models.User) error {
	res := configs.DB.Exec("INSERT INTO users(email, password_hash) VALUES (?, ?)", user.Email, user.PasswordHash)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := "SELECT id, email, password_hash, created_at FROM users WHERE email = ? LIMIT 1"

	row := configs.DB.Raw(query, email).Row()
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}
