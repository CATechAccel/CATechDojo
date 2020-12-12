package main

import (
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

	http.ListenAndServe(":8080", r)
}
