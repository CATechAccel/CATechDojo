package gacha

import (
	"CATechDojo/db"
)

type gachaInterface interface {
	SelectAllOdds() ([]CharacterData, error)
	SelectHitCharacter(HitCharacterID string) (*CharacterData, error)
	SelectUserID(token string) (string, error)
	InsertHitCharacter(token string, userCharacterID string, HitCharacterID string) error
}

func New() gachaInterface {
	return &CharacterData{}
}

func (g *CharacterData) SelectAllOdds() ([]CharacterData, error) {
	rows, err := db.DBInstance.Query("SELECT character_id, odds FROM gacha_odds")
	if err != nil {
		return nil, err
	}

	oddsSlice := make([]CharacterData, 0)
	for rows.Next() {
		var g CharacterData
		if err := rows.Scan(&g.CharacterID, &g.Odds); err != nil {
			return nil, err
		}
		oddsSlice = append(oddsSlice, g)
	}
	return oddsSlice, nil
}

func (g *CharacterData) SelectHitCharacter(HitCharacterID string) (*CharacterData, error) {
	row := db.DBInstance.QueryRow("SELECT * FROM characters WHERE id = ?", HitCharacterID)
	if err := row.Scan(&g.CharacterID, &g.Name, &g.Odds); err != nil {
		return nil, err
	}
	return g, nil
}

func (g *CharacterData) InsertHitCharacter(userCharacterID string, userID string, HitCharacterID string) error {
	if _, err := db.DBInstance.Exec("INSERT INTO user_characters(user_character_id, user_id, character_id) VALUES (?, ?, ?)", userCharacterID, userID, HitCharacterID); err != nil {
		return err
	}
	return nil
}
