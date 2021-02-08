package user_character

import "CATechDojo/db"

type userCharacterInterface interface {
	SelectUserCharactersByUserID(UserID string) ([]UserCharacterEntity, error)
}

func New() userCharacterInterface {
	return &UserCharacterEntity{}
}

func (u *UserCharacterEntity) SelectUserCharactersByUserID(UserID string) ([]UserCharacterEntity, error) {
	rows, err := db.DBInstance.Query("SELECT user_character_id, character_id FROM user_characters WHERE user_id =?", UserID)
	if err != nil {
		return nil, err
	}

	userCharacterSlice := make([]UserCharacterEntity, 0)
	for rows.Next() {
		if err := rows.Scan(&u.UserCharacterID, &u.CharacterID); err != nil {
			return nil, err
		}
		userCharacterSlice = append(userCharacterSlice, *u)
	}
	return userCharacterSlice, nil
}
