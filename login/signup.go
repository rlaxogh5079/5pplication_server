package login

import (
	"fmt"

	database "github.com/5pplication/Server/db"
)

func SignUp() {
	user, selectErr := database.SelectUser("테스트2@테스트.com")
	if selectErr == nil {
		fmt.Println("아이디가 존재합니다.")
	} else {
		fmt.Println("회원가입 시작")
	}

	fmt.Println(user)
}
