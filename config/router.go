package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mateuscapoani/go-api-crud-basic/controller"
)

func IniciarRoteamento() {
	router := gin.Default()

	controller.CriarRotas(router)

	err := router.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Erro ao inicializar roteamento")
	}
}
