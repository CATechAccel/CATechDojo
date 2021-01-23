package main

import (
<<<<<<< HEAD
	"CATechDojo/controller/character"
=======
	"CATechDojo/controller/gacha"
>>>>>>> 85773d5901a5c4e33996bc3c2e0f91e7267d7151
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
<<<<<<< HEAD
	r.HandleFunc("/character/list", character.ShowUserCharacters).Methods("GET")
=======
	r.HandleFunc("/gacha/draw", gacha.Draw).Methods("GET")
>>>>>>> 85773d5901a5c4e33996bc3c2e0f91e7267d7151

	http.ListenAndServe(":8080", r)
}
