package response

type UserCharacterResponse struct {
	UserCharacterID string `json:"user_character_id"`
	CharacterID     string `json:"character_id"`
	Name            string `json:"name"`
}

type CharactersResponse struct {
	Characters []UserCharacterResponse `json:"character"`
}
