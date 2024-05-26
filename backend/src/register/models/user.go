package models

import (
	"EDA.Project.ERP.backend.register/data"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}

func (User) GetUserInfoByUsername(username string) (User, error) {
	var user User
	err := data.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (User) AddUser(username string, password string) (int, error) {
	user := User{Username: username, Password: password}
	err := data.Db.Create(&user).Error
	return user.Id, err
}
