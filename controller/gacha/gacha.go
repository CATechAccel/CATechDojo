package gacha

import (
	"CATechDojo/controller/response"
	"CATechDojo/model/character"
	"CATechDojo/model/gacha"
	"CATechDojo/model/user"
	"CATechDojo/service/util"
	"CATechDojo/view"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func Draw(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	if token == "" {
		log.Println("トークンの値がnilです")
		http.Error(w, "認証情報が必要です。", http.StatusBadRequest)
		return
	}

	reqBody, err := view.ReadGachaDrawRequest(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//userIDの取得
	u := user.New()
	userData, err := u.SelectUserByToken(token)
	if err != nil {
		log.Println(err)
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	//gacha_oddsテーブルから全件取得
	g := gacha.New()
	odds, err := g.SelectAll()
	if err != nil {
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	var hitCharacterID string
	hitCharactersData := make([]character.CharacterEntity, 0)

	//乱数を作成するための初期化
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < reqBody.Times; i++ {
		//乱数を作成
		random := rand.Intn(oddsSum(odds))

		//当選キャラクターの決定
		var count int
		for i, _ := range odds {
			count += odds[i].Odds

			if count < random {
				continue
			}
			hitCharacterID = odds[i].CharacterID
			break
		}

		//該当キャラクターのデータを取得
		c := character.New()
		hitCharacterData, err := c.SelectCharacterByCharacterID(hitCharacterID)
		if err != nil {
			log.Println(err)
			http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
		}

		hitCharactersData = append(hitCharactersData, *hitCharacterData)

		//ユーザーキャラクターIDの作成
		userCharacterID, err := util.CreateUUID()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := c.InsertCharacterData(userCharacterID, userData.UserID, hitCharacterID); err != nil {
			log.Println(err)
			http.Error(w, "キャラクターデータを保存できませんでした", http.StatusInternalServerError)
		}
	}

	var hitCharacterSlice response.DrawResponse
	for _, hitCharacterData := range hitCharactersData {
		res := response.DrawResult{
			CharacterID: hitCharacterData.ID,
			Name:        hitCharacterData.Name,
		}
		hitCharacterSlice.Results = append(hitCharacterSlice.Results, res)
	}

	if err := view.WriteResponse(w, hitCharacterSlice); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func oddsSum(odds []gacha.OddsEntity) int {
	var sum int
	for _, o := range odds {
		sum += o.Odds
	}
	return sum
}
