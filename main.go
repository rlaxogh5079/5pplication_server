package main

/*
import (
	database "github.com/5pplication/Server/db"
)

func main() {
	database.UpdateStoreArticle("test1@test.com", "[1234567 1234568]", "1234569")
}
*/

import (
	server "github.com/5pplication/Server/server"
)

// @termsOfService  http://swagger.io/terms/
// @contact.email rlaxogh507906@gmail.com

func main() {
	server.Api().Run()
}
