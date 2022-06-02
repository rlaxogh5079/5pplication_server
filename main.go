package main

/*
import (
	database "github.com/5pplication/Server/db"
)

func main() {
	database.UpdatePassword("test1@test.com", "test2")
}
*/
import (
	server "github.com/5pplication/Server/server"
)

func main() {
	server.Api().Run()
}
