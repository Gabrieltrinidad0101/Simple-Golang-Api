package database

import structs "main/src/Structs"

func InitMigration() (ok bool) {
	db, ok := GetConnection()
	if !ok {
		return
	}
	db.AutoMigrate(&structs.User{})
	db.AutoMigrate(&structs.Student{})
	return
}
