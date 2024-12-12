package services

import (
	"errors"

	"github.com/sammervalgas/api-go-gin-clubs/datasources"
	"github.com/sammervalgas/api-go-gin-clubs/models"
)

// Retorna todos os clubes
func GetAllClubs() ([]models.Club, error) {
	var clubs []models.Club
	if err := datasources.DB.Find(&clubs).Error; err != nil {
		return nil, err
	}
	return clubs, nil
}

// Retorna um clube pelo PID
func GetClubByPID(pid string) (*models.Club, error) {
	var club models.Club
	if err := datasources.DB.Where("pid = ?", pid).First(&club).Error; err != nil {
		return nil, errors.New("club not found")
	}
	return &club, nil
}

// Cria um novo clube
func CreateClub(club *models.Club) error {
	if err := datasources.DB.Create(&club).Error; err != nil {
		return err
	}
	return nil
}

func DeleteByPID(pid string) error {
	club, err := GetClubByPID(pid)
	if err != nil {
		return err
	}

	if err := datasources.DB.Delete(&club).Error; err != nil {
		return errors.New("failed to delete club " + pid)
	}

	datasources.DB.Delete(&club, pid)
	return nil
}
