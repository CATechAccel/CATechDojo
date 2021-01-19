package gacha

import (
	"CATechDojo/db"
)

type gachaInterface interface {
	SelectAllOdds() ([]GachaData, error)
	SelectHitCharacter(HitCharacterID string) (*GachaData, error)
	InsertHitCharacter(token string, userCharacterID string, HitCharacterID string) error
}

func New() gachaInterface {
	return &GachaData{}
}

func (g *GachaData) SelectAllOdds() ([]GachaData, error) {
	rows, err := db.DBInstance.Query("SELECT character_id, odds FROM gacha_odds")
	if err != nil {
		return nil, err
	}

	oddsSlice := make([]GachaData, 0)
	for rows.Next() {
		var g GachaData
		if err := rows.Scan(&g.CharacterID, &g.Odds); err != nil {
			return nil, err
		}
		oddsSlice = append(oddsSlice, g)
	}
	return oddsSlice, nil
}

func (g *GachaData) SelectHitCharacter(HitCharacterID string) (*GachaData, error) {
	row := db.DBInstance.QueryRow("SELECT * FROM characters WHERE id = ?", HitCharacterID)
	if err := row.Scan(&g.CharacterID, &g.Name, &g.Odds); err != nil {
		return nil, err
	}
	return g, nil
}

func (g *GachaData) InsertHitCharacter(token string, userCharacterID string, HitCharacterID string) error {
	var userID string
	row := db.DBInstance.QueryRow("SELECT user_id FROM users WHERE auth_token = ?", token)
	if err := row.Scan(&userID); err != nil {
		return err
	}
	if _, err := db.DBInstance.Exec("INSERT INTO user_characters(user_character_id, user_id, character_id) VALUES (?, ?, ?)", userCharacterID, userID, HitCharacterID); err != nil {
		return err
	}
	return nil
}
