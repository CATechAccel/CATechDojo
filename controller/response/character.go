package response

import "CATechDojo/model/character"

type CharactersResponse struct {
	Characters []character.UserCharacterData `json:"character"`
}
