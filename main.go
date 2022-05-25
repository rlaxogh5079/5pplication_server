package main

import (
	"github.com/5pplication/Server/db"
	"github.com/5pplication/Server/login"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func main() {
	insertPassword, hashErr := bcrypt.GenerateFromPassword([]byte("테스트"), bcrypt.DefaultCost)
	insertErr := db.InsertUser("테스트", "테스트", "테스트", string(insertPassword))
	fmt.Println(hashErr, insertErr)
	login.Login()
}
