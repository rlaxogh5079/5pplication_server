package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

func LoadUsers(db *sql.DB) ([]map[string]interface{}, error) {
	var detectedErr error = nil

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

func InsertUser(db *sql.DB, userEmail string, nickname string, userPassword string, userStoreArticle string) (bool, error) {
	var detectedErr error = nil
	var generatedErr error = nil

	var user User
	user.Email = userEmail
	user.Nickname = nickname
	user.Password, generatedErr = bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	user.StoreArticle = userStoreArticle

	if checkErr(generatedErr) {
		detectedErr = generatedErr
	}
	result, execErr := db.Exec(fmt.Sprintf("INSERT INTO user VALUE (\"%v\", \"%v\", \"%v\", \"%v\");", user.Email, user.Nickname, string(user.Password), user.StoreArticle))
	if checkErr(execErr) {
		detectedErr = execErr
		return false, detectedErr
	}
	count, affectErr := result.RowsAffected()
	if checkErr(affectErr) {
		detectedErr = affectErr
	}
	if count != 0 {
		return true, detectedErr
	} else {
		return false, detectedErr
	}
}

func SelectUser(db *sql.DB, userEmail string) (User, error) {
	var detectedErr error = nil

	var user User

	queryErr := db.QueryRow(fmt.Sprintf("SELECT * FROM user WHERE email=\"%v\"", userEmail)).Scan(&user.Email, &user.Nickname, &user.Password, &user.StoreArticle)
	fmt.Println(queryErr)
	if checkErr(queryErr) {
		detectedErr = queryErr
	}
	fmt.Println(user)
	return user, detectedErr
}

func RemoveUser(db *sql.DB, userEmail string) (bool, error) {
	var detectedErr error = nil

	result, deleteErr := db.Exec(fmt.Sprintf("DELETE FROM user WHERE email=\"%v\"", userEmail))
	if checkErr(deleteErr) {
		detectedErr = deleteErr
		return false, detectedErr
	}

	flag, affectErr := isOne(result)
	if checkErr(affectErr) {
		detectedErr = affectErr
	}
	return flag, detectedErr
}

func UpdateNickname(db *sql.DB, userEmail string, nickname string) (bool, error) {
	var detectedErr error = nil

	result, updateErr := db.Exec(fmt.Sprintf("UPDATE user SET nickname=\"%v\" WHERE email=\"%v\"", nickname, userEmail))

	if checkErr(updateErr) {
		detectedErr = updateErr
		return false, detectedErr
	}

	flag, affectErr := isOne(result)
	if checkErr(affectErr) {
		detectedErr = affectErr
	}
	return flag, detectedErr
}

func UpdatePassword(db *sql.DB, userEmail string, userPassword string) (bool, error) {
	var detectedErr error = nil

	newPassword, generatedErr := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if checkErr(generatedErr) {
		detectedErr = generatedErr
	}

	result, updateErr := db.Exec(fmt.Sprintf("UPDATE user SET password=\"%v\" WHERE email=\"%v\"", string(newPassword), userEmail))

	if checkErr(updateErr) {
		detectedErr = updateErr
		return false, detectedErr
	}

	flag, affectErr := isOne(result)
	if checkErr(affectErr) {
		detectedErr = affectErr
	}
	return flag, detectedErr
}

func UpdateStoreArticle(db *sql.DB, userEmail string, atclNo string) (bool, error) {
	var detectedErr error = nil

	user, selectErr2 := SelectUser(db, userEmail)
	if checkErr(selectErr2) {
		detectedErr = selectErr2
	}

	storeArticle := user.StoreArticle
	if storeArticle[0:] != "" {
		articleList := append(strings.Split(strings.Trim(storeArticle, "[]"), " "), atclNo)
		storeArticle = fmt.Sprintf("%v", articleList)
	} else {
		articleList := append(strings.Split(strings.Trim(storeArticle, "[]"), ""), atclNo)
		storeArticle = fmt.Sprintf("%v", articleList)
	}

	result, updateErr := db.Exec(fmt.Sprintf("UPDATE user SET storeArticle=\"%v\" WHERE email=\"%v\"", storeArticle, userEmail))
	if checkErr(updateErr) {
		detectedErr = updateErr
		return false, detectedErr
	}

	flag, affectErr := isOne(result)
	if checkErr(affectErr) {
		detectedErr = affectErr
	}

	return flag, detectedErr
}

func RemoveStoreArticle(db *sql.DB, userEmail string, atclNo string) (bool, error) {
	var detectedErr error = nil

	user, selectErr2 := SelectUser(db, userEmail)
	if checkErr(selectErr2) {
		detectedErr = selectErr2
	}

	storeArticle := user.StoreArticle
	flag, articleList := Pop(atclNo, strings.Split(strings.Trim(storeArticle, "[]"), " "))
	if !flag {
		return false, detectedErr
	}
	storeArticle = fmt.Sprintf("%v", articleList)

	result, updateErr := db.Exec(fmt.Sprintf("UPDATE user SET storeArticle=\"%v\" WHERE email=\"%v\"", storeArticle, userEmail))
	if checkErr(updateErr) {
		detectedErr = updateErr
		return false, detectedErr
	}

	flag2, affectErr := isOne(result)
	if checkErr(affectErr) {
		detectedErr = affectErr
	}

	return flag2, detectedErr
}
