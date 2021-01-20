package characters

import (
	"CATechDojo/controller/response"
	"CATechDojo/model/characters"
	"encoding/json"
	"log"
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	if token == "" {
		log.Println("トークンの値がnilです")
		http.Error(w, "認証情報が必要です。", http.StatusBadRequest)
		return
	}

	c := characters.New()

	//user_idの取得
	userID, err := c.SelectUserID(token)
	if err != nil {
		log.Println(err)
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	//user_idを用いてuser_character_id, character_idを取得
	userCharacters, err := c.SelectUserCharacters(userID)
	if err != nil {
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	var userCharacterSlice []characters.UserCharacterData
	for _, u := range userCharacters {
		characterName, err := c.SelectCharacterName(u.CharacterID)
		if err != nil {
			log.Println(err)
			http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
		}
		res := characters.UserCharacterData{
			UserCharacterID: u.UserCharacterID,
			CharacterID:     u.CharacterID,
			Name:            characterName,
		}
		userCharacterSlice = append(userCharacterSlice, res)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	res := response.CharactersResponse{userCharacterSlice}
	data, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(data)

}
