package user

import (
	"CATechDojo/db"
)

//インターフェースを定義
type userInterface interface {
	SelectUser(string) (*UserData, error)
	InsertUser(UserData)
	UpdateUser(UserData) (UserData, error)
}

//定義したインターフェースを満たすインスタンスを生成する関数を定義
func New() userInterface {
	return &UserData{}
}

//インスタンスが持つ関数（メソッド）を定義
func (u *UserData) SelectUser(token string) (*UserData, error) {
	row := db.DBInstance.QueryRow("SELECT * FROM user WHERE auth_token = ?", token)
	if err := row.Scan(&u.UserID, &u.AuthToken, &u.Name); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserData) InsertUser(data UserData) {
	panic("implement me")
}

func (u *UserData) UpdateUser(data UserData) (UserData, error) {
	panic("implement me")
}
