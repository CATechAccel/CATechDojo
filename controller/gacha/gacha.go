package gacha

import (
	"CATechDojo/model/gacha"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Draw(w http.ResponseWriter, r *http.Request) {
	//x-tokenを受け取る
	token := r.Header.Get("x-token")
	if token == "" {
		log.Println("トークンの値がnilです")
		http.Error(w, "認証情報が必要です。", http.StatusBadRequest)
		return
	}

	//timesを受け取る
	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var times int
	if err := json.Unmarshal(buf.Bytes(), &times); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//特定のキャラクターをDBから取り出す
	c := gacha.New()

	if err := c.Select(); err != nil {
		log.Println(err)
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	//取り出したキャラクターをDBに保存する
	if err := c.Insert(); err != nil {
		log.Println(err)
		http.Error(w, "ユーザデータを保存できませんでした", http.StatusInternalServerError)
	}

	//取り出したキャラクターのidとnameをjson形式で返す
}
