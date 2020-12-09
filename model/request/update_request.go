package request

import "CATechDojo/db"

type UpdateRequest struct {
	Name string `json:"name"`
}

func (u *UpdateRequest) UpdateName(token string) error {
	if _, err := db.DBInstance.Exec("UPDATE user SET name = ? WHERE auth_token = ?", u.Name, token); err != nil {
		return err
	}
	return nil
}
