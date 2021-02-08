package user

import (
	"CATechDojo/controller/request"
	"CATechDojo/controller/response"
	"CATechDojo/model/user"
	"CATechDojo/service/util"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	u := user.New()

	users, err := u.SelectAll()
	if err != nil {
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var userSlice response.GetAllUserRespponse
	for _, userdata := range users {
		res := response.UserResponse{
			UserID: userdata.UserID,
			Name:   userdata.Name,
		}
		userSlice.Users = append(userSlice.Users, res)
	}

	data, err := json.Marshal(userSlice)
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
	userData, err := u.SelectUserByToken(token)
	if err != nil {
		log.Println(err)
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	res := response.GetUserResponse{
		Name: userData.Name,
	}

	data, err := json.Marshal(res)
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

	var reqBody request.UserRequest
	if err := json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	userID, err := util.CreateUUID()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	authToken, err := util.CreateUUID()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	newUser := user.UserEntity{
		UserID:    userID,
		AuthToken: authToken,
		Name:      reqBody.Name,
	}

	if err := newUser.Insert(); err != nil {
		log.Println(err)
		http.Error(w, "ユーザデータを保存できませんでした", http.StatusInternalServerError)
	}

	var res response.CreateUserResponse
	res.Token = newUser.AuthToken

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(res)
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

	var reqBody request.UserRequest
	if err := json.Unmarshal(buf.Bytes(), &reqBody); err != nil {
		errorResponse(err, w)
	}

	var u user.UserEntity
	u.Name = reqBody.Name
	if err := u.UpdateName(token); err != nil {
		errorResponse(err, w)
	}
}

func errorResponse(err error, w http.ResponseWriter) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
