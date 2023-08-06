package repository

import (
	"fmt"

	"github.com/mateuscapoani/go-api-crud-basic/db"
	"github.com/mateuscapoani/go-api-crud-basic/model"
)

func GetLivro(idLivro int) *model.Livro {
	var livro model.Livro
	db.DB.First(&livro, idLivro)
	return &livro
}

func CriarLivro(livro *model.Livro) {
	resultado := db.DB.Create(&livro)

	if resultado.Error != nil {
		fmt.Println("Erro ao criar objeto")
	}
}
