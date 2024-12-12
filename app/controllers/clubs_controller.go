package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sammervalgas/api-go-gin-clubs/models"
	"github.com/sammervalgas/api-go-gin-clubs/services"
)

func ListAllClubs(c *gin.Context) {
	clubs, err := services.GetAllClubs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clubs)
}

func FindByPID(c *gin.Context) {
	pid := c.Params.ByName("pid")

	club, err := services.GetClubByPID(pid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, club)
}

func CreateClub(c *gin.Context) {
	var club models.Club

	if err := c.ShouldBindJSON(&club); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateClub(&club); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, club)
}

func DeleteByPID(c *gin.Context) {
	pid := c.Query("pid")

	if err := services.DeleteByPID(pid); err != nil {
		var errorm = models.ErrorMessage{
			Code:    http.StatusNotFound,
			Message: "PID not found",
		}
		c.JSON(http.StatusNotFound, errorm)
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})

}
