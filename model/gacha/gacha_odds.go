package gacha

import (
	"CATechDojo/db"
)

type gachaInterface interface {
	SelectAll() ([]OddsData, error)
}

func New() gachaInterface {
	return &OddsData{}
}

func (g *OddsData) SelectAll() ([]OddsData, error) {
	rows, err := db.DBInstance.Query("SELECT * FROM gacha_odds")
	if err != nil {
		return nil, err
	}

	oddsSlice := make([]OddsData, 0)
	for rows.Next() {
		var g OddsData
		if err := rows.Scan(&g.CharacterID, &g.Odds); err != nil {
			return nil, err
		}
		oddsSlice = append(oddsSlice, g)
	}
	return oddsSlice, nil
}
