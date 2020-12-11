package user

import (
	"CATechDojo/controller/request"
	"CATechDojo/model/user"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	u := user.New()

	users, err := u.SelectAll()
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

func GetUser(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	data, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(data)
}

func Create(w http.ResponseWriter, r *http.Request) {
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

	userID, err := createUUID()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	authToken, err := createUUID()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	reqBody.UserID = userID
	reqBody.AuthToken = authToken

	if err := reqBody.Insert(); err != nil {
		log.Println(err)
		http.Error(w, "ユーザデータを保存できませんでした", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(reqBody)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(data)
}

func ChangeName(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")

	body := r.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		errorResponse(err, w)
	}

	var reqBody request.UpdateNameRequest
	if err := json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		errorResponse(err, w)
	}

	var u user.UserData
	u.Name = reqBody.Name
	if err := u.UpdateName(token); err != nil {
		errorResponse(err, w)
	}
}

func errorResponse(err error, w http.ResponseWriter) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func createUUID() (string, error) {
	uuID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uuID.String(), nil
}
