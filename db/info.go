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
	Nickname     string `json:"name"`
	Password     []byte `json:"password"`
	StoreArticle string `json:"storeArticle"`
}

type UserWithoutPassword struct {
	Email        string `json:"email"`
	Nickname     string `json:"name"`
	StoreArticle string `json:"storeArticle"`
}

type Article struct {
	AtclNo  string                 `json:"atclNo"`
	Email   string                 `json:"email"`
	Share   bool                   `json:"share"`
	Long    float32                `json:"longitude"`
	Lat     float32                `json:"latitude"`
	Title   string                 `json:"title"`
	Body    string                 `json:"body"`
	Date    string                 `json:"date"`
	Likecnt int                    `json:"likecnt"`
	Tag     map[string]interface{} `json:"tag"`
}

func checkErr(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}

func isOne(result sql.Result) (bool, error) {
	count, affectErr := result.RowsAffected()
	if count == 1 {
		return true, affectErr
	} else {
		return false, affectErr
	}
}

func ConnectDB() (*sql.DB, error) {
	db, mysqlErr := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4", Username, Password, Host, Port, Database))
	if !checkErr(mysqlErr) {
		fmt.Println("데이터베이스 연결 성공")
	}

	return db, mysqlErr
}
