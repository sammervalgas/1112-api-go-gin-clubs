package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sammervalgas/api-go-gin-clubs/controllers"
)

func HandleRequest() {

	r := gin.Default()

	r.GET("/clubs", controllers.ListAllClubs)

	r.GET("/clubs/:pid", controllers.FindByPID)

	r.POST("/clubs", controllers.CreateClub)

	// User routes
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetUserById)
	r.POST("/users", controllers.CreateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.PATCH("/users/:id", controllers.UpdateUser)

	r.Run()
}
