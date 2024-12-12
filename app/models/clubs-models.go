package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Club struct {
	gorm.Model

	Pid uuid.UUID `json:"pid"`

	Name string `json:"name"`

	Address string `json:"address"`

	Website string `json:"website"`
}
