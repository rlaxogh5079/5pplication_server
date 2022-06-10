package database

import (
	"encoding/json"
	"fmt"
	"strings"

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

func InsertUser(userEmail string, nickname string, userPassword string, userStoreArticle string) (bool, error) {
	var detectedErr error = nil
	var generatedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

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

func SelectUser(userEmail string) (User, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

	var user User

	rows, dataErr := db.Query(fmt.Sprintf("SELECT *  FROM user WHERE email=\"%v\"", userEmail))
	fmt.Println(dataErr)
	if checkErr(dataErr) {
		detectedErr = dataErr
	}
	for rows.Next() {
		loadErr := rows.Scan(&user.Email, &user.Password, &user.Nickname, &user.StoreArticle)
		if checkErr(loadErr) {
			detectedErr = loadErr
		}
	}
	fmt.Println(user)
	return user, detectedErr
}

func RemoveUser(userEmail string) (bool, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

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

func UpdateNickname(userEmail string, nickname string) (bool, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

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

func UpdatePassword(userEmail string, userPassword string) (bool, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

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

func UpdateStoreArticle(userEmail string, atclNo string) (bool, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

	user, selectErr2 := SelectUser(userEmail)
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

func RemoveStoreArticle(userEmail string, atclNo string) (bool, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

	user, selectErr2 := SelectUser(userEmail)
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
