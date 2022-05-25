package main

import (
	"github.com/5pplication/Server/rest"
)

/*
func main() {
	insertPassword, hashErr := bcrypt.GenerateFromPassword([]byte("테스트"), bcrypt.DefaultCost)
	insertErr := db.InsertUser("테스트", "테스트", "테스트", string(insertPassword))
	fmt.Println(hashErr, insertErr)
	login.Login()
}
*/

func main() {
	rest.Api().Run()
}
