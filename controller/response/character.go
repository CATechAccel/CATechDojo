package response

import "CATechDojo/model/characters"

type CharactersResponse struct {
	Characters []characters.UserCharacterData `json:"characters"`
}
