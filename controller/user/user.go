package user

import (
	"CATechDojo/model/user"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	u := user.New()

	users, err := u.SelectAllUser()
	if err != nil {
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	data, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(data)
}

func ChangeUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")

	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var reqBody user.UserData
	if err := json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := reqBody.UpdateUser(token); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
