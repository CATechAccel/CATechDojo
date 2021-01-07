package gacha

import (
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
	//取り出したキャラクターのidとnameをjson形式で返す
}
