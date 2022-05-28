package login

import (
	"fmt"

	database "github.com/5pplication/Server/db"
	"golang.org/x/crypto/bcrypt"
)

func CheckLogin(userEmail string, userPassword []byte) (int, error) { // 로그인 성공 -> 1, 로그인 실패 -> 0, 아이디 존재 X -> -1
	user, selectErr := database.SelectUser(userEmail)

	if user.Email == userEmail {
		compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), userPassword)
		if compareErr != nil {
			return 0, selectErr
		} else {
			return 1, selectErr
		}
	}
	return -1, selectErr
}

func Login(userEmail string, userPassword []byte) bool {
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
