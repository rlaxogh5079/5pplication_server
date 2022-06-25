package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func LoadArticle(db *sql.DB, long float64, lat float64) ([]map[string]interface{}, error) {
	// 데이터베이스로 부터 해당 좌표에 있는 모든 글을 불러오는 함수
	var detectedErr error = nil

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

func InsertArticle(db *sql.DB, atclNo string, email string, share bool, long string, lat string, title string, body string, likecnt int64, date string, image string, tag string) (bool, error) {
	// 데이터베이스에 글 정보를 입력하는 함수
	var detectedErr error = nil

	var article = &Article{atclNo, email, share, long, lat, title, body, date, image, likecnt, tag}

	result, insertErr := db.Exec(fmt.Sprintf("INSERT INTO article VALUE (%v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v);", article.AtclNo, article.Email, article.Share, article.Long, article.Lat, article.Title, article.Body, article.Likecnt, article.Date, article.Images, article.Tag))
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

func SelectArticle(db *sql.DB, atclNo string) (Article, error) {
	// 해당 글번호를 가진 글을 불러옴
	var detectedErr error = nil

	var article Article

	queryErr := db.QueryRow(fmt.Sprintf("SELECT * FROM article WHERE atclNo=\"%v\"", atclNo)).Scan(&article.AtclNo, &article.Email, &article.Share, &article.Long, &article.Lat, &article.Title, &article.Body, &article.Likecnt, &article.Date, &article.Images, &article.Tag)
	if checkErr(queryErr) {
		detectedErr = queryErr
	}
	return article, detectedErr
}

func SelectUserArticle(db *sql.DB, userEmail string) (map[string]interface{}, error) {
	var detectedErr error = nil

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

func RemoveArticle(db *sql.DB, atclNo string) (bool, error) {
	var detectedErr error = nil

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
