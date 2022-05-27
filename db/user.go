package database

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func LoadUsers() ([]map[string]interface{}, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
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

		loadErr := rows.Scan(&user.Email, &user.Nickname, &user.Password, &user.StoreArticle)
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

func InsertUser(userEmail string, userName string, userPassword string, userStoreArticle string) error {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

	var user User
	user.Email = userEmail
	user.Nickname = userName
	user.Password = userPassword
	user.StoreArticle = userStoreArticle

	statement, prepareErr := db.Prepare("INSERT INTO user VALUE (?, ?, ?, ?);")
	if checkErr(prepareErr) {
		detectedErr = prepareErr
	}
	_, insertErr := statement.Exec(user.Email, user.Nickname, user.Password, userStoreArticle)
	if checkErr(insertErr) {
		fmt.Printf("(%s)데이터가 이미 존재함\n", user.Email) // UNIQUE KEY
		detectedErr = mysqlErr
	} else {
		fmt.Println("데이터 삽입성공")
	}
	return detectedErr
}
