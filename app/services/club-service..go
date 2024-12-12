package services

import (
	"errors"

	"github.com/sammervalgas/api-go-gin-clubs/datasources"
	"github.com/sammervalgas/api-go-gin-clubs/models"
)

// GetAllClubs retorna uma lista com todos os clubes cadastrados no banco de dados.
// Retorna um slice de Club e um possível erro caso ocorra falha na consulta.
func GetAllClubs() ([]models.Club, error) {
	var clubs []models.Club
	if err := datasources.DB.Find(&clubs).Error; err != nil {
		return nil, err
	}
	return clubs, nil
}

// GetClubByPID busca um clube específico pelo seu PID (identificador único).
// Recebe o PID como string e retorna um ponteiro para Club e possível erro.
// Retorna erro "club not found" caso o clube não seja encontrado.
func GetClubByPID(pid string) (*models.Club, error) {
	var club models.Club
	if err := datasources.DB.Where("pid = ?", pid).First(&club).Error; err != nil {
		return nil, errors.New("club not found")
	}
	return &club, nil
}

// CreateClub persiste um novo clube no banco de dados.
// Recebe um ponteiro para Club e retorna erro em caso de falha na criação.
// O clube deve conter todos os campos obrigatórios preenchidos.
func CreateClub(club *models.Club) error {
	if err := datasources.DB.Create(&club).Error; err != nil {
		return err
	}
	return nil
}

// DeleteByPID remove um clube do banco de dados usando seu PID.
// Primeiro verifica se o clube existe, depois realiza a deleção.
// Retorna erro caso o clube não seja encontrado ou ocorra falha na deleção.
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
