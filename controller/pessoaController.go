package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mateuscapoani/go-api-crud-basic/service"
)

func CriarRotas(router *gin.Engine) {
	router.GET("/:idPessoa", getPessoa)
}

func getPessoa(c *gin.Context) {
	result := service.GetPessoa(c.Param("idPessoa"))
	c.JSON(200, gin.H{
		"message": result,
	})
}
