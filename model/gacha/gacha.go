package gacha

import "CATechDojo/db"

type gachaInterface interface {
	Select() error
	Insert(userChracterID string, token string) error
}

func New() gachaInterface {
	return &CharacterData{}
}

func (c *CharacterData) Select() error {
	row := db.DBInstance.QueryRow("SELECT id, name FROM characters WHERE id = ?", "character_0001")
	if err := row.Scan(&c.ID, &c.Name); err != nil {
		return err
	}
	return nil
}

func (c *CharacterData) Insert(userCharacterID string, token string) error {
	var userID string
	row := db.DBInstance.QueryRow("SELECT user_id FROM users WHERE auth_token = ?", token)
	if err := row.Scan(&userID); err != nil {
		return err
	}
	if _, err := db.DBInstance.Exec("INSERT INTO user_characters(user_character_id, user_id, character_id) VALUES (?, ?, ?)", userCharacterID, userID, c.ID); err != nil {
		return err
	}
	return nil
}
