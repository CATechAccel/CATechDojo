package user

import (
	"CATechDojo/db"
)

//インターフェースを定義
type userInterface interface {
	SelectAll() ([]UserEntity, error)
	SelectUserByToken(token string) (*UserEntity, error)
	Insert() error
	UpdateName(token string) error
}

func New() userInterface {
	return &UserEntity{}
}

func (u *UserEntity) SelectAll() ([]UserEntity, error) {
	rows, err := db.DBInstance.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	userSlice := make([]UserEntity, 0)
	for rows.Next() {
		var u UserEntity
		if err := rows.Scan(&u.UserID, &u.AuthToken, &u.Name); err != nil {
			return nil, err
		}
		userSlice = append(userSlice, u)
	}
	return userSlice, nil
}

func (u *UserEntity) SelectUserByToken(token string) (*UserEntity, error) {
	row := db.DBInstance.QueryRow("SELECT * FROM users WHERE auth_token = ?", token)
	if err := row.Scan(&u.UserID, &u.AuthToken, &u.Name); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserEntity) Insert() error {
	if _, err := db.DBInstance.Exec("INSERT INTO users(user_id, auth_token, name) VALUES (?, ?, ?)", u.UserID, u.AuthToken, u.Name); err != nil {
		return err
	}
	return nil
}

func (u *UserEntity) UpdateName(token string) error {
	if _, err := db.DBInstance.Exec("UPDATE users SET name = ? WHERE auth_token = ?", u.Name, token); err != nil {
		return err
	}
	return nil
}
