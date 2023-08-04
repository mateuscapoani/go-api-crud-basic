package config

import (
	"log"

	"github.com/joho/godotenv"
)

func CarregarVariaveisDeAmbiente() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao inicializar variáveis de ambiente")
	}
}
