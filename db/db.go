package db

import (
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CriarConexao() {
	var err error

	var dsn strings.Builder

	dsn.WriteString("host=")
	dsn.WriteString(os.Getenv("DBHOST"))
	dsn.WriteString(" user=")
	dsn.WriteString(os.Getenv("DBUSER"))
	dsn.WriteString(" password=")
	dsn.WriteString(os.Getenv("DBPASSWORD"))
	dsn.WriteString(" dbname=")
	dsn.WriteString(os.Getenv("DBNAME"))
	dsn.WriteString(" port=")
	dsn.WriteString(os.Getenv("DBPORT"))
	dsn.WriteString(" sslmode=disable")

	DB, err = gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao criar conexao com banco de dados")
	}
}
