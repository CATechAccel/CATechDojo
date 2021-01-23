package gacha

import "CATechDojo/db"

func (g *CharacterData) SelectUserID(token string) (string, error) {
	var userID string
	row := db.DBInstance.QueryRow("SELECT user_id FROM users WHERE auth_token = ?", token)
	if err := row.Scan(&userID); err != nil {
		return "", err
	}
	return userID, nil
}
