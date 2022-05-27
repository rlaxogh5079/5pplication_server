package login

import (
	"fmt"

	database "github.com/5pplication/Server/db"
	"golang.org/x/crypto/bcrypt"
)

func CheckLogin(userEmail string, userPassword []byte) (int, error) { // 로그인 성공 -> 1, 로그인 실패 -> 0, 아이디 존재 X -> -1
	var user database.User
	var detectedErr error = nil
	db, mysqlErr := database.ConnectDB()
	if mysqlErr != nil {
		detectedErr = mysqlErr
		panic(mysqlErr)
	}
	defer db.Close()

	dataErr := db.QueryRow(fmt.Sprintf("SELECT * FROM user WHERE email=\"%v\"", userEmail)).Scan(&user.Email, &user.Nickname, &user.Password, &user.StoreArticle) // 테이블에서 정보를 가져와 rows에 저장
	if dataErr != nil {
		detectedErr = dataErr
		panic(dataErr)
	} else {
		fmt.Println("데이터 로드 성공")
	}

	if user.Email == userEmail {
		compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), userPassword)
		if compareErr != nil {
			return 0, detectedErr
		} else {
			return 1, detectedErr
		}
	}
	return -1, detectedErr
}
func Login() bool {
	userEmail := "테스트" // Flutter로 부터 입력
	userPassword, hashErr := bcrypt.GenerateFromPassword([]byte("테스트"), bcrypt.DefaultCost)
	if hashErr != nil {
		panic(hashErr)
	}
	flag, _ := CheckLogin(userEmail, userPassword)
	if flag == 1 {
		fmt.Println("로그인 성공")
		return true
	} else if flag == 0 {
		fmt.Println("로그인 실패")
	} else {
		fmt.Println("회원가입")
	}
	return false
}
