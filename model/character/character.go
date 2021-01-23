package character

import "CATechDojo/db"

type characterInterface interface {
	SelectUserCharacters(UserID string) ([]UserCharacterData, error)
	SelectCharacterName(CharacterID string) (string, error)
}

func New() characterInterface {
	return &CharacterData{}
}

func (c *CharacterData) SelectUserCharacters(UserID string) ([]UserCharacterData, error) {
	rows, err := db.DBInstance.Query("SELECT user_character_id, character_id FROM user_characters WHERE user_id =?", UserID)
	if err != nil {
		return nil, err
	}

	userCharacterSlice := make([]UserCharacterData, 0)
	for rows.Next() {
		var u UserCharacterData
		if err := rows.Scan(&u.UserCharacterID, &u.CharacterID); err != nil {
			return nil, err
		}
		userCharacterSlice = append(userCharacterSlice, u)
	}
	return userCharacterSlice, nil
}

func (c *CharacterData) SelectCharacterName(CharacterID string) (string, error) {
	var name string
	row := db.DBInstance.QueryRow("SELECT name FROM characters WHERE id = ?", CharacterID)
	if err := row.Scan(&name); err != nil {
		return "", err
	}
	return name, nil
}

/*

func SelectUserID(token string) (string, error) {
	var userID string
	row := db.DBInstance.QueryRow("SELECT user_id FROM users WHERE auth_token = ?", token)
	if err := row.Scan(&userID); err != nil {
		return "", err
	}
	return userID, nil
}

func SelectUserCharacters(UserID string) ([]UserCharacterData, error) {
	rows, err := db.DBInstance.Query("SELECT user_character_id, character_id FROM user_characters WHERE user_id =?", UserID)
	if err != nil {
		return nil, err
	}

	userCharacterSlice := make([]UserCharacterData, 0)
	for rows.Next() {
		var u UserCharacterData
		if err := rows.Scan(&u.UserCharacterID, &u.CharacterID); err != nil {
			return nil, err
		}
		userCharacterSlice = append(userCharacterSlice, u)
	}
	return userCharacterSlice, nil
}

func SelectCharacterName(CharacterID string) (string, error) {
	var c UserCharacterData
	row := db.DBInstance.QueryRow("SELECT name FROM characters WHERE id = ?", CharacterID)
	if err := row.Scan(&c.Name); err != nil {
		return "", err
	}
	return c.Name, nil
}

*/
