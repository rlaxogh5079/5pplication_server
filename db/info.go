package database

import (
	"database/sql"
	"fmt"
)

const (
	Host     = "132.226.16.134"
	Database = "5pplicationDB"
	Username = "5pplication"
	Password = "5pplication"
	Port     = "3306"
)

type User struct {
	Email        string `json:"email"`
	Id           string `json:"id"`
	Nickname     string `json:"name"`
	Password     string `json:"password"`
	StoreArticle string `json:"storeArticle"`
}

type Article struct {
	AtclNo   string                 `json:"atclNo"`
	UserId   string                 `json:"userId"`
	Long     float32                `json:"longitude"`
	Lat      float32                `json:"latitude"`
	Title    string                 `json:"title"`
	Body     string                 `json:"body"`
	Date     string                 `json:"date"`
	Likecnt  int                    `json:"likecnt"`
	Comments map[string]interface{} `json:"comments"`
}


func checkErr(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}

func ConnectDB() (*sql.DB, error) {
	db, mysqlErr := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", Username, Password, Host, Port, Database))
	if !checkErr(mysqlErr) {
		fmt.Println("데이터베이스 연결 성공")
	}

	return db, mysqlErr
}
