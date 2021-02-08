package character

import (
	"CATechDojo/db"
)

// モデル層＝DBとのアクセスのみ関心をもつ（今回のプロジェクトでは）
type characterInterface interface {
	SelectCharacterByCharacterID(CharacterID string) (*CharacterEntity, error)
	InsertCharacterData(userCharacterID string, userID string, HitCharacterID string) error
	GetName() string
}

func New() characterInterface {
	return &CharacterEntity{}
}

func (c *CharacterEntity) GetName() string {
	return c.Name
}

func (c *CharacterEntity) SelectCharacterByCharacterID(CharacterID string) (*CharacterEntity, error) {
	row := db.DBInstance.QueryRow("SELECT id, name, power FROM characters WHERE id = ?", CharacterID)
	if err := row.Scan(&c.ID, &c.Name, &c.Power); err != nil {
		return nil, err
	}
	return c, nil
}

func (g *CharacterEntity) InsertCharacterData(userCharacterID string, userID string, CharacterID string) error {
	if _, err := db.DBInstance.Exec("INSERT INTO user_characters(user_character_id, user_id, character_id) VALUES (?, ?, ?)", userCharacterID, userID, CharacterID); err != nil {
		return err
	}
	return nil
}
