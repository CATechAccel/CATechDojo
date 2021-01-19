package response

type DrawResponse struct {
	CharacterID string `json:"character_id"`
	Name        string `json:"name"`
}

type DrawAllResponse struct {
	Results []DrawResponse `json:"results"`
}
