package response

type DrawResult struct {
	CharacterID string `json:"character_id"`
	Name        string `json:"name"`
}

type DrawResponse struct {
	Results []DrawResult `json:"results"`
}
