package user

import (
	"CATechDojo/controller/response"
	"CATechDojo/model/user"
	"CATechDojo/service/util"
	"CATechDojo/view"
	"encoding/json"
	"log"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	u := user.New()
	users, err := u.SelectAll()
	if err != nil {
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}

	var userSlice response.GetAllUserRespponse
	for _, userdata := range users {
		res := response.UserResponse{
			UserID: userdata.UserID,
			Name:   userdata.Name,
		}
		userSlice.Users = append(userSlice.Users, res)
	}

	if err := view.WriteResponse(w, userSlice); err != nil {
		log.Println(err)
		http.Error(w, "データを参照できませんでした", http.StatusInternalServerError)
	}
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

	if err := view.WriteResponse(w, res); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = w.Write(data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	reqBody, err := view.ReadUserRequest(r)
	if err != nil {
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
	if err := view.WriteResponse(w, res); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ChangeName(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	if token == "" {
		log.Println("トークンの値がnilです")
		http.Error(w, "認証情報が必要です。", http.StatusBadRequest)
		return
	}

	reqBody, err := view.ReadUserRequest(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var u user.UserEntity
	u.AuthToken = token
	if err := u.UpdateName(reqBody.Name); err != nil {
		errorResponse(err, w)
	}
}

func errorResponse(err error, w http.ResponseWriter) {
	log.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
