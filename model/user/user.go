package user

import (
	"CATechDojo/db"
)

//インターフェースを定義
type userInterface interface {
	SelectAll() ([]UserData, error)
	SelectUser(token string) (*UserData, error)
	Insert() error
	UpdateName(token string) error
	GetName() string
	GetUserID() string
}

func New() userInterface {
	return &UserData{}
}

func (u *UserData) GetName() string {
	return u.Name
}

func (u *UserData) GetUserID() string {
	return u.UserID
}

func (u *UserData) SelectAll() ([]UserData, error) {
	rows, err := db.DBInstance.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	userSlice := make([]UserData, 0)
	for rows.Next() {
		var u UserData
		if err := rows.Scan(&u.UserID, &u.AuthToken, &u.Name); err != nil {
			return nil, err
		}
		userSlice = append(userSlice, u)
	}
	return userSlice, nil
}

// TODO: SelectUserByTokenみたいなわかりやすい名前をつける
func (u *UserData) SelectUser(token string) (*UserData, error) {
	row := db.DBInstance.QueryRow("SELECT * FROM users WHERE auth_token = ?", token)
	if err := row.Scan(&u.UserID, &u.AuthToken, &u.Name); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserData) Insert() error {
	if _, err := db.DBInstance.Exec("INSERT INTO users(user_id, auth_token, name) VALUES (?, ?, ?)", u.UserID, u.AuthToken, u.Name); err != nil {
		return err
	}
	return nil
}

func (u *UserData) UpdateName(token string) error {
	if _, err := db.DBInstance.Exec("UPDATE users SET name = ? WHERE auth_token = ?", u.Name, token); err != nil {
		return err
	}
	return nil
}
