package config

import (
	"github.com/gin-gonic/gin"
	"github.com/mateuscapoani/go-api-crud-basic/controller"
)

func IniciarRoteamento() {
	router := gin.Default()

	controller.CriarRotas(router)

	router.Run()
}
