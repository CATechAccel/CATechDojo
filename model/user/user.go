package user

import "CATechDojo/db"

type UserData struct {
	UserID    string `json:"user_id"`
	AuthToken string `json:"auth_token"`
	Name      string `json:"name"`
}

func SelectAllUser() ([]UserData, error) {
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
