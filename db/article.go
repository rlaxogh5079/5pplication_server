package database

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func LoadArticle(long float64, lat float64) ([]map[string]interface{}, error) {
	// 데이터베이스로 부터 해당 좌표에 있는 모든 글을 불러오는 함수
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

	rows, dataErr := db.Query(fmt.Sprintf("SELECT atclNo, email, longitude, latitude, title, body, date FROM article WHERE longitude=\"%v\" AND latitude=\"%v\"", long, lat))
	if checkErr(dataErr) {
		detectedErr = dataErr
	} else {
		fmt.Println("데이터 로드 성공")
	}
	defer rows.Close()

	var articles []map[string]interface{}

	for rows.Next() {
		var article Article
		var articleData map[string]interface{}
		loadErr := rows.Scan(&article.AtclNo, &article.Email, &article.Long, &article.Lat, &article.Title, &article.Body, &article.Date)
		if checkErr(loadErr) {
			detectedErr = loadErr
		}

		articleByte, _ := json.Marshal(article)
		json.Unmarshal(articleByte, &articleData)
		fmt.Printf("데이터 : %v\n", articleData)
		articles = append(articles, articleData)
	}
	if detectedErr == nil {
		fmt.Println("데이터 반환 완료")
	}
	return articles, detectedErr
}

func InsertArticle(atclNo string, email string, share bool, long string, lat string, title string, body string, likecnt int64, date string, imagesStr string, tagStr string) (bool, error) {
	// 데이터베이스에 글 정보를 입력하는 함수
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

	var imageData map[string]interface{}
	var tagData map[string]interface{}

	images, _ := json.Marshal(imagesStr)
	json.Unmarshal(images, &imageData)

	tags, _ := json.Marshal(tagStr)
	json.Unmarshal(tags, &tagData)

	var article Article
	article.AtclNo = atclNo
	article.Email = email
	article.Share = share
	article.Long = long
	article.Lat = lat
	article.Title = title
	article.Body = body
	article.Date = date
	article.Images = imageData
	article.Likecnt = likecnt
	article.Tag = tagData

	statement, prepareErr := db.Prepare("INSERT INTO article VALUE (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if checkErr(prepareErr) {
		detectedErr = prepareErr
	}
	result, insertErr := statement.Exec(article.AtclNo, article.Email, article.Share, article.Long, article.Lat, article.Title, article.Body, article.Likecnt, article.Date, article.Images, article.Tag)
	if checkErr(insertErr) {
		detectedErr = insertErr
		return false, insertErr
	}

	count, affectErr := result.RowsAffected()
	if checkErr(affectErr) {
		detectedErr = affectErr
	}

	return bool(count != 0), detectedErr
}

func SelectArticle(atclNo string) (Article, error) {
	// 해당 글번호를 가진 글을 불러옴
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

	var article Article

	queryErr := db.QueryRow(fmt.Sprintf("SELECT * FROM article WHERE atclNo=\"%v\"", atclNo)).Scan(&article.AtclNo, &article.Email, &article.Share, &article.Long, &article.Lat, &article.Title, &article.Body, &article.Likecnt, &article.Date, &article.Tag)
	if checkErr(queryErr) {
		detectedErr = queryErr
	}
	return article, detectedErr
}

func SelectUserArticle(userEmail string) (map[string]interface{}, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

	rows, dataErr := db.Query(fmt.Sprintf("SELECT * FROM article WHERE email=\"%s\"", userEmail)) // userEmail이 작성한 글의 내용을 전부 조회
	if checkErr(dataErr) {
		detectedErr = dataErr
	}

	var user User
	var userData map[string]interface{}
	for rows.Next() {
		loadErr := rows.Scan(&user.Email, &user.Nickname, &user.Password, &user.StoreArticle)
		if checkErr(loadErr) {
			detectedErr = loadErr
		}

		fmt.Printf("데이터 : %v", user)
		userByte, _ := json.Marshal(user)
		json.Unmarshal(userByte, &userData)
	}

	return userData, detectedErr
}

func RemoveArticle(atclNo string) (bool, error) {
	var detectedErr error = nil
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		detectedErr = mysqlErr
	}
	defer db.Close()

	result, deleteErr := db.Exec(fmt.Sprintf("DELETE FROM article WHERE atclno=\"%v\"", atclNo))
	if deleteErr != nil {
		detectedErr = deleteErr
	}

	count, affectErr := result.RowsAffected()
	if affectErr != nil {
		detectedErr = affectErr
	}
	if count == 0 {
		fmt.Println("데이터가 삭제되지 않았습니다.")
		return false, detectedErr
	} else {
		fmt.Printf("%v가 삭제되었습니다.\n", atclNo)
		return true, detectedErr
	}
}
