package gacha

type CharacterData struct {
	CharacterID string `json:"character_id"`
	Name        string `json:"name"`
	Odds        int    `json:"odds"`
}