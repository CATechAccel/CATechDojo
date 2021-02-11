package gacha

import (
	"CATechDojo/db"
)

type gachaInterface interface {
	SelectAll() ([]OddsEntity, error)
}

func New() gachaInterface {
	return &OddsEntity{}
}

func (g *OddsEntity) SelectAll() ([]OddsEntity, error) {
	rows, err := db.DBInstance.Query("SELECT * FROM gacha_odds")
	if err != nil {
		return nil, err
	}

	oddsSlice := make([]OddsEntity, 0)
	for rows.Next() {
		var g OddsEntity
		if err := rows.Scan(&g.CharacterID, &g.Odds); err != nil {
			return nil, err
		}
		oddsSlice = append(oddsSlice, g)
	}
	return oddsSlice, nil
}
