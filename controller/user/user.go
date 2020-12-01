package user

import (
	"CATechDojo/model/user"
	"encoding/json"
	"net/http"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := user.SelectAllUser()
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
