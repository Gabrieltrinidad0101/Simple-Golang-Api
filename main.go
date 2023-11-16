package main

import (
	server "main/src"
	database "main/src/Database"
)

func main() {
	database.InitMigration()
	server.StartServer()
}
