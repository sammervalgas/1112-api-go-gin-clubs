package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sammervalgas/api-go-gin-clubs/datasources/entities"
	"github.com/sammervalgas/api-go-gin-clubs/services"
)

// CreateUser cria um novo usuário com os dados recebidos no corpo da requisição
// Retorna status 201 Created com os dados do usuário criado em caso de sucesso
// ou status 400 Bad Request se os dados forem inválidos
func CreateUser(c *gin.Context) {
	var user entities.User
	// Faz o binding do JSON recebido para a struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	// Cria o usuário através do service e retorna os dados
	user = services.CreateUser(user)
	c.JSON(http.StatusCreated, user)
}

// GetAllUsers retorna uma lista de todos os usuários cadastrados
// Retorna status 200 OK com a lista de usuários
func GetAllUsers(c *gin.Context) {
	// Busca todos os usuários através do service
	users := services.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// GetUserById busca um usuário específico pelo seu ID
// Retorna status 200 OK com os dados do usuário em caso de sucesso
// ou status 404 Not Found se o usuário não for encontrado
// Caso especial: Se o nome do usuário for "joao", retorna uma mensagem especial
func GetUserById(c *gin.Context) {
	// Extrai o ID dos parâmetros da URL
	id := c.Params.ByName("id")
	user := services.GetUserById(id)

	// Verifica se o usuário existe
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found"})
		return
	}

	// Caso especial para usuário com nome "joao"
	if user.Name == "joao" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Usuario premiado",
			"user":    user,
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser remove um usuário pelo seu ID
// Retorna status 204 No Content em caso de sucesso na remoção
// ou status 404 Not Found se o usuário não for encontrado
func DeleteUser(c *gin.Context) {
	// Extrai o ID dos parâmetros da URL
	id := c.Params.ByName("id")
	user := services.GetUserById(id)

	// Verifica se o usuário existe
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found"})
		return
	}

	// Remove o usuário e retorna status 204
	services.DeleteUser(id)
	c.JSON(http.StatusNoContent, nil)
}

// UpdateUser atualiza os dados de um usuário existente
// Retorna status 200 OK com os dados atualizados em caso de sucesso
// ou status 400 Bad Request se os dados forem inválidos
// ou status 404 Not Found se o usuário não for encontrado
func UpdateUser(c *gin.Context) {
	// Extrai o ID dos parâmetros da URL
	id := c.Params.ByName("id")
	var user entities.User

	// Faz o binding do JSON recebido para a struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	// Verifica se o usuário existe
	existingUser := services.GetUserById(id)
	if existingUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found"})
		return
	}

	// Atualiza os dados do usuário e retorna
	updatedUser := services.UpdateUser(id, user)
	c.JSON(http.StatusOK, updatedUser)
}
