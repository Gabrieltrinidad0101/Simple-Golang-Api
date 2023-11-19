package database

import (
	"fmt"
	"main/src/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, bool) {
	conf := utils.Configuration{}
	conf.LoadEnviroments()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DbUser,
		conf.DbPassword,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)
	db, error := gorm.Open(mysql.Open(dsn))
	if error != nil {
		fmt.Printf("ERROR IN THE CONNECTION \n %s", error)
		return nil, false
	}
	return db, true
}
