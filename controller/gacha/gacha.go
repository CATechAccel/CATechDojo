package gacha

import (
	"CATechDojo/controller/user"
	"CATechDojo/model/gacha"
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
	c := gacha.New()

	if err := c.SelectCharacter(); err != nil {
		log.Println(err)
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	//取り出したキャラクターをDBに保存する
	userCharacterID, err := user.CreateUUID()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := c.InsertCharacter(userCharacterID, token); err != nil {
		log.Println(err)
		http.Error(w, "ユーザデータを保存できませんでした", http.StatusInternalServerError)
	}

	//取り出したキャラクターのidとnameをjson形式で返す
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(c)
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

	var times struct{}
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

	//乱数を作成
	rand.Seed(time.Now().UnixNano())
	randam := rand.Intn(oddsSum(odds))

	//カウントアップ用変数を定義
	var count int
	var hitCharacterID *string

	//当選キャラクターの決定
	for i, _ := range odds {
		count += odds[i].Odds

		if count < randam {
			continue
		}
		h := odds[i].CharacterID
		hitCharacterID = &h
		break
	}

	if err := g.SelectHitCharacter(*hitCharacterID); err != nil {
		log.Println()
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	userCharacterID, err := user.CreateUUID()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := g.InsertHitCharacter(token, userCharacterID, *hitCharacterID); err != nil {
		log.Println(err)
		http.Error(w, "ユーザデータを保存できませんでした", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(g)
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
