package main

import database "github.com/5pplication/Server/db"

func main() {
	database.RemoveUser("test4@test.com")
}
