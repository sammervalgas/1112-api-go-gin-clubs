package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string    `json:"name" gorm:"type:varchar(100);not null"`
	CPF      string    `json:"cpf" gorm:"type:varchar(11);unique;not null"`
	Birthday time.Time `json:"birthday" gorm:"not null"`
}
