package services

import (
	"github.com/sammervalgas/api-go-gin-clubs/datasources"
	"github.com/sammervalgas/api-go-gin-clubs/datasources/entities"
)

func CreateUser(user entities.User) entities.User {
	datasources.DB.Create(&user)
	return user
}

func GetAllUsers() []entities.User {
	var users []entities.User
	datasources.DB.Find(&users)
	return users
}

func GetUserById(id string) entities.User {
	var user entities.User
	datasources.DB.First(&user, id)
	return user
}

func DeleteUser(id string) {
	var user entities.User
	datasources.DB.Delete(&user, id)
}

func UpdateUser(id string, user entities.User) entities.User {
	var existingUser entities.User
	datasources.DB.First(&existingUser, id)
	datasources.DB.Model(&existingUser).UpdateColumns(user)
	return existingUser
}
