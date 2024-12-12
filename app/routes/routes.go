package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sammervalgas/api-go-gin-clubs/controllers"
)

// HandleRequest configura todas as rotas da API e inicia o servidor
func HandleRequest() {
	// Cria uma nova instância do router Gin com configuração padrão
	r := gin.Default()

	// Rotas para gerenciamento de Clubes
	// Lista todos os clubes cadastrados
	r.GET("/clubs", controllers.ListAllClubs)

	// Busca um clube específico pelo PID (identificador único)
	r.GET("/clubs/:pid", controllers.FindByPID)

	// Cria um novo clube
	r.POST("/clubs", controllers.CreateClub)

	// Rotas para gerenciamento de Usuários
	// Lista todos os usuários cadastrados
	r.GET("/users", controllers.GetAllUsers)

	// Busca um usuário específico pelo ID
	r.GET("/users/:id", controllers.GetUserById)

	// Cria um novo usuário
	r.POST("/users", controllers.CreateUser)

	// Remove um usuário pelo ID
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Atualiza os dados de um usuário existente
	r.PATCH("/users/:id", controllers.UpdateUser)

	// Inicia o servidor na porta padrão :8080
	r.Run()
}
