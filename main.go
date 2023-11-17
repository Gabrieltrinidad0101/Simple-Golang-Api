package main

import (
	server "main/src"
	database "main/src/Database"
)

func main() {
	ok := database.InitMigration()
	if !ok {
		return
	}
	server.StartServer()
}
