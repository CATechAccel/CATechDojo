package db

//参考：learn-golang 8.データベース
//https://github.com/CATechAccel/learn-golang/pull/20/files

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//MySQLと接続するためのライブラリ
	_ "github.com/go-sql-driver/mysql"
	//.envを利用するためのライブラリ
	"github.com/joho/godotenv"
)

var DBInstance *sql.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file: %v", err)
	}

	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbhost := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")

	//user:password@tcp(host:port)/dbname
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", dbuser, dbpassword, dbhost, dbname)

	var err error
	DBInstance, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	log.Printf("データベースへの接続に成功しました;%s\n", dataSource)
}
