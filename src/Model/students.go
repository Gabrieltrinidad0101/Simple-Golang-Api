package model

import (
	database "main/src/Database"
	structs "main/src/Structs"
)

type Student struct{}

func NewStudent() *Student {
	return &Student{}
}

func (u *Student) Create(student *structs.Student) (ok bool) {
	db, ok := database.GetConnection()
	if !ok {
		return
	}
	db.Create(&student)
	return
}

func (u *Student) Get() (students []structs.Student, ok bool) {
	db, ok := database.GetConnection()
	if !ok {
		return
	}
	db.Find(&students)
	return
}

func (u *Student) Update(student structs.Student) (ok bool) {
	db, ok := database.GetConnection()
	if !ok {
		return
	}
	db.Model(&student).Updates(student)
	return
}

func (u *Student) Delete(studentId uint) (ok bool) {
	db, ok := database.GetConnection()
	if !ok {
		return
	}
	db.Delete(&structs.User{}, studentId)
	return
}
