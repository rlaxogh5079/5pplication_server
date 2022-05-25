package database

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func LoadComment(userId string) []map[string]interface{} {
	db, mysqlErr := ConnectDB()
	if checkErr(mysqlErr) {
		panic(mysqlErr)
	}
	defer db.Close()

	rows, dataErr := db.Query(fmt.Sprintf("SELECT * FROM comment WHERE userId=\"%s\"", userId))
	checkErr(dataErr)
	defer rows.Close()
	fmt.Println("Successfully Load Data!")

	var comment Comment
	var comments []map[string]interface{}
	var commentData map[string]interface{}

	for rows.Next() { // 한줄씩 데이터를 읽어들임

		loadErr := rows.Scan(&comment.CommentNo, &comment.UserId, &comment.Dislikecnt, &comment.Likecnt, &comment.Body, &comment.Commentcnt)
		checkErr(loadErr)

		// JSON 형식으로 변환
		commentByte, _ := json.Marshal(comment)
		json.Unmarshal(commentByte, &commentData)
		fmt.Printf("Data : %v\n", commentData)
		comments = append(comments, commentData)
	}
	return comments
}
