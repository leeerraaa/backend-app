package psql

import (
	"database/sql"
	"errors"

	"github.com/leeerraaa/backend-app/internal/domain"
)

type Auth struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *Auth {
	return &Auth{db: db}
}

func (a *Auth) UserInfo(userId string) (domain.User, error) {
	rows, err := a.db.Query("SELECT * FROM users WHERE id = $1", userId)
	if err != nil {
		return domain.User{}, err
	}

	var userData domain.User
	for rows.Next() {
		if err := rows.Scan(&userData.Id, &userData.Login, &userData.Username, &userData.Password, &userData.DateOfCreation); err != nil {
			return domain.User{}, err
		}
	}

	return userData, rows.Err()
}

func (a *Auth) GetUser(login, password string) (string, error) {
	var userId string
	rows, err := a.db.Query("SELECT id FROM users WHERE login = $1 AND password = $2", login, password)
	if err != nil {
		return "", err
	}

	for rows.Next() {
		if err := rows.Scan(&userId); err != nil {
			return "", err
		}
	}

	if userId == "" {
		return "", errors.New("userId not found")
	}

	return userId, rows.Err()
}
