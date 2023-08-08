package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mateuscapoani/go-api-crud-basic/dto/request"
	"github.com/mateuscapoani/go-api-crud-basic/service"
)

func criarRotasLivros(router *gin.Engine) {
	grupoLivro := router.Group("/livro")
	grupoLivro.GET("/:idLivro", getLivro)
	grupoLivro.POST("", criarLivro)
}

func getLivro(context *gin.Context) {
	idLivro, err := strconv.Atoi(context.Param("idLivro"))
	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	result, err := service.GetLivro(idLivro)
	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(200, result)
}

func criarLivro(context *gin.Context) {
	var livro request.LivroRequest
	context.Bind(&livro)

	idLivro, err := service.CriarLivro(&livro)
	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(201, idLivro)
}
