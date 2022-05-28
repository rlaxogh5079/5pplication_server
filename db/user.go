package database

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

func LoadUsers() ([]map[string]interface{}, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

	rows, dataErr := db.Query("SELECT email, nickname, storeArticle FROM user")
	if checkErr(dataErr) {
		detectedErr = dataErr
	} else {
		fmt.Println("데이터 로드 성공")
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() {

		var user UserWithoutPassword
		var userData map[string]interface{}

		loadErr := rows.Scan(&user.Email, &user.Nickname, &user.StoreArticle)
		if checkErr(loadErr) {
			detectedErr = loadErr
		}

		userByte, _ := json.Marshal(user)
		json.Unmarshal(userByte, &userData)
		fmt.Printf("데이터 : %v\n", userData)
		users = append(users, userData)
	}
	if detectedErr == nil {
		fmt.Println("데이터 반환 완료")
	}
	return users, detectedErr
}

func InsertUser(userEmail string, userName string, userPassword string, userStoreArticle string) error {
	var detectedErr error = nil
	var generatedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

	var user User
	user.Email = userEmail
	user.Nickname = userName
	user.Password, generatedErr = bcrypt.GenerateFromPassword([]byte(userPassword),bcrypt.DefaultCost)
	user.StoreArticle = userStoreArticle

	if checkErr(generatedErr) {
		detectedErr = generatedErr
	}
	statement, prepareErr := db.Prepare("INSERT INTO user VALUE (?, ?, ?, ?);")
	if checkErr(prepareErr) {
		detectedErr = prepareErr
	}
	_, insertErr := statement.Exec(user.Email, user.Nickname, user.Password, userStoreArticle)
	if checkErr(insertErr) {
		fmt.Printf("(%s)데이터가 이미 존재함\n", user.Email) // Primary Key
		detectedErr = insertErr
	} else {
		fmt.Println("데이터 삽입성공")
	}
	return detectedErr
}

func SelectUser(userEmail string) (User, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

	var user User

	queryErr := db.QueryRow(fmt.Sprintf("SELECT * FROM user WHERE email=\"%v\"", userEmail)).Scan(&user.Email, &user.Nickname, &user.Password, &user.StoreArticle)

	if checkErr(queryErr) {
		detectedErr = queryErr
	}

	return user, detectedErr
}
/*
func RemoveUser(userEmail string) (bool, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()


}
*/