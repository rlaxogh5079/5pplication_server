package login

import (
	"fmt"
	"strings"
	"github.com/5pplication/Server/db"
	"golang.org/x/crypto/bcrypt"
)

func Login() bool {
	userId := "테스트" // Flutter로 부터 입력
	userPassword, hashErr := bcrypt.GenerateFromPassword([]byte("테스트"), bcrypt.DefaultCost)
	if hashErr != nil {
		panic(hashErr)
	}
	if strings.Contains(userId,"@") {
		flag,_ := db.CheckLogin(userId, userPassword, "")
		if flag == 1 {
			fmt.Println("로그인 성공")
			return true
		} else if flag == 0 {
			fmt.Println("로그인 실패")
		} else {
			fmt.Println("회원가입")
		}
	} else {
		flag,_ := db.CheckLogin(userId, userPassword, "")
		if flag == 1 {
			fmt.Println("로그인 성공")
			return true
		} else if flag == 0 {
			fmt.Println("로그인 실패")
		} else {
			fmt.Println("회원가입")
		}
		
	}
	return false
}