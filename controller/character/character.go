package character

import (
	"CATechDojo/controller/response"
	"CATechDojo/model/character"
	"CATechDojo/model/user"
	"encoding/json"
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
	if err := u.SelectUser(token); err != nil {
		log.Println(err)
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	// user_idを用いてuser_character_id, character_idを取得
	c := character.New()
	userCharacters, err := c.SelectUserCharacters(u.GetUserID())
	if err != nil {
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	var userCharacterResponseSlice []response.UserCharacterResponse
	for _, userCharacter := range userCharacters {
		characterName, err := c.SelectCharacterName(userCharacter.CharacterID)
		if err != nil {
			log.Println(err)
			http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
		}
		res := response.UserCharacterResponse{
			UserCharacterID: userCharacter.UserCharacterID,
			CharacterID:     userCharacter.CharacterID,
			Name:            characterName,
		}
		userCharacterResponseSlice = append(userCharacterResponseSlice, res)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	res := response.CharactersResponse{userCharacterResponseSlice}
	data, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(data)

}

/*

func ShowUserCharacters(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	if token == "" {
		log.Println("トークンの値がnilです")
		http.Error(w, "認証情報が必要です。", http.StatusBadRequest)
		return
	}

	var userCharacterSlice []character.UserCharacterData
	for _, u := range getUserCharactersData(token) {
		userCharacterData := character.UserCharacterData{
			UserCharacterID: u.UserCharacterID,
			CharacterID:     u.CharacterID,
			Name:            getCharacterName(u.CharacterID),
		}
		userCharacterSlice = append(userCharacterSlice, userCharacterData)
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

func getUserID(token string) string {
	userID, err := character.SelectUserID(token)
	if err != nil {
		fmt.Println("データを参照できませんでした")
	}
	return userID
}

func getUserCharactersData(token string) []character.UserCharacterData {
	userCharacters, err := character.SelectUserCharacters(getUserID(token))
	if err != nil {
		fmt.Println("データを参照できませんでした")
	}
	return userCharacters
}

func getCharacterName(characterID string) string {
	characterName, err := character.SelectCharacterName(characterID)
	if err != nil {
		fmt.Println("データを参照できませんでした")
	}
	return characterName
}

*/
