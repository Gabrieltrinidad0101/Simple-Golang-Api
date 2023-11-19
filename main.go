package main

import (
	server "main/src"
	database "main/src/Database"
)

/*
TO RUN THE TESTS IT IS NECESSARY TO ADD THE .ENV FILE IN THE CONTROLLERS FOLDER WITH THE TEST ONES
RUN TEST WITH go test ./src/Controllers
*/
func main() {
	ok := database.InitMigration()
	if !ok {
		return
	}
	server.StartServer()
}
