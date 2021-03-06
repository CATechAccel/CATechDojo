package character

import (
	"CATechDojo/controller/response"
	"CATechDojo/model/character"
	"CATechDojo/model/user"
	"CATechDojo/model/user_character"
	"CATechDojo/view"
	"log"
	"net/http"
)

func ShowUserCharacters(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	if token == "" {
		log.Println("トークンの値がnilです")
		http.Error(w, "認証情報が必要です。", http.StatusBadRequest)
		return
	}

	u := user.New()
	userData, err := u.SelectUserByToken(token)
	if err != nil {
		log.Println(err)
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	// user_idを用いてuser_character_id, character_idを取得
	uc := user_character.New()
	userCharacters, err := uc.SelectUserCharactersByUserID(userData.UserID)
	if err != nil {
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	c := character.New()
	var userCharacterResponseSlice []response.UserCharacterResponse
	for _, userCharacter := range userCharacters {
		characterData, err := c.SelectCharacterByCharacterID(userCharacter.CharacterID)
		if err != nil {
			log.Println(err)
			http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
		}
		res := response.UserCharacterResponse{
			UserCharacterID: userCharacter.UserCharacterID,
			CharacterID:     userCharacter.CharacterID,
			Name:            characterData.Name,
		}
		userCharacterResponseSlice = append(userCharacterResponseSlice, res)
	}

	if err := view.WriteResponse(w, userCharacterResponseSlice); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
