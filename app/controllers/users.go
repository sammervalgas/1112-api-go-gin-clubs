package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sammervalgas/api-go-gin-clubs/datasources/entities"
	"github.com/sammervalgas/api-go-gin-clubs/services"
)

func CreateUser(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	user = services.CreateUser(user)
	c.JSON(http.StatusCreated, user)
}

func GetAllUsers(c *gin.Context) {
	users := services.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	user := services.GetUserById(id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found"})
		return
	}

	if user.Name == "joao" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Usuario premiado",
			"user":    user,
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user := services.GetUserById(id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found"})
		return
	}

	services.DeleteUser(id)
	c.JSON(http.StatusNoContent, nil)
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user entities.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	existingUser := services.GetUserById(id)
	if existingUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found"})
		return
	}

	updatedUser := services.UpdateUser(id, user)
	c.JSON(http.StatusOK, updatedUser)
}
