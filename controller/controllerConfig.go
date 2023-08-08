package controller

import "github.com/gin-gonic/gin"

func CriarRotas(router *gin.Engine) {
	criarRotasLivros(router)
}
