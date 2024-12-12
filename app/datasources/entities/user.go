package entities

import (
	"time"

	"gorm.io/gorm"
)

// User representa a entidade de usuário no sistema
// Contém informações básicas como nome, CPF e data de nascimento
type User struct {
	gorm.Model // Incorpora os campos padrão do GORM (ID, CreatedAt, UpdatedAt, DeletedAt)

	// Nome do usuário com limite de 100 caracteres
	Name string `json:"name" gorm:"type:varchar(100);not null"`

	// CPF do usuário - deve ser único e ter exatamente 11 dígitos
	CPF string `json:"cpf" gorm:"type:varchar(11);unique;not null"`

	// Data de nascimento do usuário
	Birthday time.Time `json:"birthday" gorm:"not null"`
}
