package datasources

import (
	"log"

	"github.com/sammervalgas/api-go-gin-clubs/datasources/entities"
	"github.com/sammervalgas/api-go-gin-clubs/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnection() {
	dadosConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dadosConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Club{}, &entities.User{})
}
