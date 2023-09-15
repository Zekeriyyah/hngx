package models

import (
	"github.com/Zekeriyyah/hngx/stagetwop/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name    string `gorm:"" json:"name"`
	Email   string `json:"email"`
	Country string `json:"country"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (p *User) CreateUser() *User {
	db.NewRecord(p)
	db.Create(&p)
	return p
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var user User
	db := db.Where("ID = ?", Id).Find(&user)
	return &user, db
}

func DeleteUser(Id int64) User {
	var user User
	db.Where("ID = ?", Id).Delete(user)
	return user
}
