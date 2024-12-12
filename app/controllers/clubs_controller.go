package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sammervalgas/api-go-gin-clubs/models"
	"github.com/sammervalgas/api-go-gin-clubs/services"
)

// ListAllClubs retorna uma lista de todos os clubes cadastrados
// Retorna status 200 OK com a lista de clubes em caso de sucesso
// ou status 500 Internal Server Error em caso de falha
func ListAllClubs(c *gin.Context) {
	// Busca todos os clubes através do service
	clubs, err := services.GetAllClubs()
	if err != nil {
		// Em caso de erro, retorna status 500 com a mensagem de erro
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Retorna a lista de clubes com status 200
	c.JSON(http.StatusOK, clubs)
}

// FindByPID busca um clube específico pelo seu identificador único (PID)
// Retorna status 200 OK com os dados do clube em caso de sucesso
// ou status 404 Not Found se o clube não for encontrado
func FindByPID(c *gin.Context) {
	// Extrai o PID dos parâmetros da URL
	pid := c.Params.ByName("pid")

	// Busca o clube pelo PID através do service
	club, err := services.GetClubByPID(pid)
	if err != nil {
		// Se o clube não for encontrado, retorna status 404
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Retorna os dados do clube encontrado com status 200
	c.JSON(http.StatusOK, club)
}

// CreateClub cria um novo clube com os dados recebidos no corpo da requisição
// Retorna status 201 Created com os dados do clube criado em caso de sucesso
// ou status 400 Bad Request se os dados forem inválidos
// ou status 500 Internal Server Error em caso de falha na criação
func CreateClub(c *gin.Context) {
	var club models.Club

	// Faz o binding do JSON recebido para a struct Club
	if err := c.ShouldBindJSON(&club); err != nil {
		// Se os dados forem inválidos, retorna status 400
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tenta criar o clube através do service
	if err := services.CreateClub(&club); err != nil {
		// Em caso de erro na criação, retorna status 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Retorna o clube criado com status 201
	c.JSON(http.StatusCreated, club)
}

// DeleteByPID remove um clube pelo seu identificador único (PID)
// Retorna status 204 No Content em caso de sucesso na remoção
// ou status 404 Not Found se o clube não for encontrado
func DeleteByPID(c *gin.Context) {
	// Obtém o PID da query string da URL
	pid := c.Query("pid")

	// Tenta deletar o clube através do service
	if err := services.DeleteByPID(pid); err != nil {
		// Se o clube não for encontrado, cria uma mensagem de erro personalizada
		var errorm = models.ErrorMessage{
			Code:    http.StatusNotFound,
			Message: "PID not found",
		}
		// Retorna status 404 com a mensagem de erro
		c.JSON(http.StatusNotFound, errorm)
		return
	}

	// Retorna status 204 sem conteúdo para indicar sucesso na deleção
	c.JSON(http.StatusNoContent, gin.H{})

}
