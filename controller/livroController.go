package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mateuscapoani/go-api-crud-basic/model"
	"github.com/mateuscapoani/go-api-crud-basic/service"
)

func CriarRotas(router *gin.Engine) {
	router.GET("/livro/:idLivro", getLivro)
	router.POST("/livro", criarLivro)
}

func getLivro(context *gin.Context) {
	idLivro, err := strconv.Atoi(context.Param("idLivro"))
	if err != nil {
		context.JSON(400, gin.H{
			"message": "Parametro informado incorreto",
		})
		return
	}

	result := service.GetLivro(idLivro)
	context.JSON(200, result)
}

func criarLivro(context *gin.Context) {
	var body struct {
		Nome    string
		Autor   string
		Editora string
	}

	context.Bind(&body)

	livro := model.Livro{Nome: body.Nome, Autor: body.Autor, Editora: body.Editora}
	service.CriarLivro(&livro)

	context.JSON(201, livro)
}
