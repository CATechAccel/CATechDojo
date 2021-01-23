package user_character

import "CATechDojo/db"

type userCharacterInterface interface {
	SelectUserCharacters(UserID string) ([]UserCharacterData, error)
}

func New() userCharacterInterface {
	return &UserCharacterData{}
}

func (u *UserCharacterData) SelectUserCharacters(UserID string) ([]UserCharacterData, error) {
	rows, err := db.DBInstance.Query("SELECT user_character_id, character_id FROM user_characters WHERE user_id =?", UserID)
	if err != nil {
		return nil, err
	}

	userCharacterSlice := make([]UserCharacterData, 0)
	for rows.Next() {
		if err := rows.Scan(&u.UserCharacterID, &u.CharacterID); err != nil {
			return nil, err
		}
		userCharacterSlice = append(userCharacterSlice, *u)
	}
	return userCharacterSlice, nil
}
