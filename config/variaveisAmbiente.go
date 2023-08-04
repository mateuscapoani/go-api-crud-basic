package config

import (
	"log"

	"github.com/joho/godotenv"
)

func CarregarVariaveisDeAmbiente() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao criar conexao com banco de dados")
	}
}
