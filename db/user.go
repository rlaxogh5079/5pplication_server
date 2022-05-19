package db

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func LoadUsers() ([]map[string]interface{}, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr){
		detectedErr = mysqlErr
	}
	defer db.Close()

	rows, dataErr := db.Query("SELECT * FROM user") // 테이블에서 정보를 가져와 rows에 저장
	if checkErr(dataErr) {
		detectedErr = dataErr
	} else {
		fmt.Println("데이터 로드 성공")
	}
	defer rows.Close()

	var users []map[string]interface{}
	for rows.Next() { // 한줄씩 데이터를 읽어들임

		var user User
		var userData map[string]interface{}

		loadErr := rows.Scan(&user.Email, &user.Id, &user.Nickname, &user.Password)
		if checkErr(loadErr) {
			detectedErr = loadErr
		}

		// JSON 형식으로 변환
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

func InsertUser(userEmail string,userId string, userName string, userPassword string) error {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr){
		detectedErr = mysqlErr
	}
	defer db.Close()

	var user User
	user.Email = userEmail
	user.Id = userId
	user.Nickname = userName
	user.Password = userPassword

	statement, prepareErr := db.Prepare(" INSERT INTO user(email, id, nickname, password) VALUE (?, ?, ?, ?);")
	if checkErr(prepareErr) {
		detectedErr = prepareErr
	}
	_, insertErr := statement.Exec(user.Email, user.Id, user.Nickname, user.Password)
	if checkErr(insertErr) {
		fmt.Printf("(%s, %s, %s)데이터가 이미 존재함\n", user.Email, user.Id, user.Nickname) // UNIQUE KEY
		detectedErr = mysqlErr
	} else {
		fmt.Println("데이터 삽입성공")
	}
	return detectedErr
}

func CheckLogin(userId string, userPassword string, userEmail string) (bool, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr){
		detectedErr = mysqlErr
	}
	defer db.Close()

	var user User
	dataErr := db.QueryRow(fmt.Sprintf("SELECT * FROM user WHERE id=\"%v\" OR email=\"%v\"",userId, userEmail)).Scan(&user.Email, &user.Id, &user.Nickname, &user.Password) // 테이블에서 정보를 가져와 rows에 저장
	if checkErr(dataErr) {
		detectedErr = dataErr
	} else {
		fmt.Println("데이터 로드 성공")
	}
	
	if userEmail=="" {
		return bool(user.Id == userId && user.Password == userPassword), detectedErr
	} else {
		return bool(user.Email == userEmail && user.Password == userPassword), detectedErr
	}
}