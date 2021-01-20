package character

import (
	"CATechDojo/controller/response"
	"CATechDojo/model/character"
	"encoding/json"
	"fmt"
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
