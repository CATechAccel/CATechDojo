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
	r.HandleFunc("/user", user.GetAllUser).Methods("GET")

	http.ListenAndServe(":8080", r)
}