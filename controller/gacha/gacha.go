package gacha

import (
	"CATechDojo/controller/request"
	"CATechDojo/controller/response"
	"CATechDojo/model/gacha"
	"CATechDojo/service/util"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func DrawSpecificCharacter(w http.ResponseWriter, r *http.Request) {
	//x-tokenを受け取る
	token := r.Header.Get("x-token")
	if token == "" {
		log.Println("トークンの値がnilです")
		http.Error(w, "認証情報が必要です。", http.StatusBadRequest)
		return
	}

	//特定のキャラクターをDBから取り出す
	g := gacha.New()

	if err := g.SelectCharacter(); err != nil {
		log.Println(err)
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	//取り出したキャラクターをDBに保存する
	userCharacterID, err := util.CreateUUID()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := g.InsertCharacter(userCharacterID, token); err != nil {
		log.Println(err)
		http.Error(w, "ユーザデータを保存できませんでした", http.StatusInternalServerError)
	}

	//取り出したキャラクターのidとnameをjson形式で返す
	res := g.GetCharacterData()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(data)
}

func Draw(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	if token == "" {
		log.Println("トークンの値がnilです")
		http.Error(w, "認証情報が必要です。", http.StatusBadRequest)
		return
	}

	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var times request.Times
	if err := json.Unmarshal(buf.Bytes(), &times); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//gacha_oddsテーブルから全件取得
	g := gacha.New()

	odds, err := g.SelectAllOdds()
	if err != nil {
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	var hitCharacterID string
	hitCharactersData := make([]gacha.GachaData, 0)

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= times.Times; i++ {
		//乱数を作成
		random := rand.Intn(oddsSum(odds))

		var count int
		//当選キャラクターの決定
		for i, _ := range odds {
			count += odds[i].Odds

			if count < random {
				continue
			}
			h := odds[i].CharacterID
			hitCharacterID = h
			break
		}

		hitCharacterData, err := g.SelectHitCharacter(hitCharacterID)
		if err != nil {
			log.Println(err)
			http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
		}

		hitCharactersData = append(hitCharactersData, *hitCharacterData)

		userCharacterID, err := util.CreateUUID()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := g.InsertHitCharacter(token, userCharacterID, hitCharacterID); err != nil {
			log.Println(err)
			http.Error(w, "ユーザデータを保存できませんでした", http.StatusInternalServerError)
		}
	}

	var hitCharacterslice response.DrawAllResponse
	for _, hitCharacterData := range hitCharactersData {
		res := response.DrawResponse{
			CharacterID: hitCharacterData.CharacterID,
			Name:        hitCharacterData.Name,
		}
		hitCharacterslice.Results = append(hitCharacterslice.Results, res)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(hitCharacterslice)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(data)
}

func oddsSum(odds []gacha.GachaData) int {
	var sum int
	for _, gachadata := range odds {
		sum += gachadata.Odds
	}
	return sum
}
