package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, bool) {
	dsn := "root:1234@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn))
	if error != nil {
		fmt.Printf("ERROR IN THE CONNECTION \n %s", error)
		return nil, false
	}
	return db, true
}
