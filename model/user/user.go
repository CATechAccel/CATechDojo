package user

import (
	"CATechDojo/db"
)

//インターフェースを定義
type userInterface interface {
	SelectAllUser() ([]UserData, error)
	UpdateUser(string) error
	InsertUser() error
}

//定義したインターフェースを満たすインスタンスを生成する関数を定義
func New() userInterface {
	return &UserData{}
}

//インスタンスが持つ関数（メソッド）を定義
func (u *UserData) SelectAllUser() ([]UserData, error) {
	rows, err := db.DBInstance.Query("SELECT * FROM user")
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

func (u *UserData) InsertUser() error {
	if _, err := db.DBInstance.Exec("INSERT INTO user(user_id, auth_token, name) VALUES (?, ?, ?)", u.UserID, u.AuthToken, u.Name); err != nil {
		return err
	}
	return nil
}

func (u *UserData) UpdateUser(token string) error {
	if _, err := db.DBInstance.Exec("UPDATE user SET name = ? WHERE auth_token = ?", u.Name, token); err != nil {
		return err
	}
	return nil
}
