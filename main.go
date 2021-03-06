package main

import (
	"CATechDojo/controller/character"
	"CATechDojo/controller/gacha"
	"CATechDojo/controller/health"
	"CATechDojo/controller/user"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//NewRouter関数で*Router型の変数を定義
	r := mux.NewRouter()

	//MethodsメソッドでHTTPメソッドの種類を指定
	r.HandleFunc("/health", health.HealthCheck).Methods("GET")
	r.HandleFunc("/user", user.GetAll).Methods("GET")
	r.HandleFunc("/user/get", user.GetUser).Methods("GET")
	r.HandleFunc("/user/create", user.Create).Methods("POST")
	r.HandleFunc("/user/update", user.ChangeName).Methods("PUT")
	r.HandleFunc("/character/list", character.ShowUserCharacters).Methods("GET")
	r.HandleFunc("/gacha/draw", gacha.Draw).Methods("GET")

	http.ListenAndServe(":8080", r)
}
