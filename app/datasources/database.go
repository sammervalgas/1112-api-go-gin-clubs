package datasources

import (
	"log"

	"github.com/sammervalgas/api-go-gin-clubs/datasources/entities"
	"github.com/sammervalgas/api-go-gin-clubs/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Variáveis globais para conexão com o banco de dados
var (
	// DB é a instância de conexão com o banco de dados
	DB *gorm.DB
	// err armazena possíveis erros de conexão
	err error
)

// DbConnection estabelece a conexão com o banco de dados PostgreSQL
// e realiza a migração automática das tabelas
func DbConnection() {
	// String de conexão com o banco de dados
	// Contém as credenciais e configurações necessárias
	dadosConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	// Abre a conexão com o banco usando o driver PostgreSQL
	DB, err = gorm.Open(postgres.Open(dadosConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	// Realiza a migração automática das tabelas
	// Cria ou atualiza as tabelas Club e User no banco
	DB.AutoMigrate(&models.Club{}, &entities.User{})
}
