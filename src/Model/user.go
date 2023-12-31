package model

import (
	database "main/src/Database"
	structs "main/src/Structs"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) Register(user structs.User) (structs.User, bool) {
	db, ok := database.GetConnection()
	if !ok {
		return structs.User{}, ok
	}
	db.Create(&user)
	return user, ok
}

func (u *User) Find(name string, password string) (user structs.User, ok bool) {
	db, ok := database.GetConnection()
	if !ok {
		return
	}
	db.First(&user, "name = ? and password = ? ", name, password)
	return
}
