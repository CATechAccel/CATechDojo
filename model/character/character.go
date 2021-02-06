package character

import (
	"CATechDojo/db"
)

type characterInterface interface {
	SelectByCharacterID(CharacterID string) (*CharacterData, error)
	InsertHitCharacterData(userCharacterID string, userID string, HitCharacterID string) error
	GetName() string
}

func New() characterInterface {
	return &CharacterData{}
}

func (c *CharacterData) GetName() string {
	return c.Name
}

func (c *CharacterData) SelectByCharacterID(CharacterID string) (*CharacterData, error) {
	row := db.DBInstance.QueryRow("SELECT id, name, power FROM characters WHERE id = ?", CharacterID)
	if err := row.Scan(&c.ID, &c.Name, &c.Power); err != nil {
		return nil, err
	}
	return c, nil
}

func (g *CharacterData) InsertHitCharacterData(userCharacterID string, userID string, HitCharacterID string) error {
	if _, err := db.DBInstance.Exec("INSERT INTO user_characters(user_character_id, user_id, character_id) VALUES (?, ?, ?)", userCharacterID, userID, HitCharacterID); err != nil {
		return err
	}
	return nil
}
