package login

import (
	"fmt"
	"strings"
	"github.com/5pplication/Server/db"
)

func Login() {
	for i := 0; i < 3; i++ {
		userId := "테스트2@gmail.com" // Flutter로 부터 입력
		userPassword := "테스트2"
		if strings.Contains(userId,"@") {
			if flag, _ := db.CheckLogin("", userPassword, userId); flag {
				fmt.Println("로그인 성공")
				break
			} else {
				fmt.Println("로그인 실패")
			}
		} else {
			if flag,_ := db.CheckLogin(userId, userPassword, ""); flag {
				fmt.Println("로그인 성공")
				break
			} else {
				fmt.Println("로그인 실패")
			}
			
		}
	}
}