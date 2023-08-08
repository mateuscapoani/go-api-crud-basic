package repository

import (
	"github.com/mateuscapoani/go-api-crud-basic/db"
	"github.com/mateuscapoani/go-api-crud-basic/model"
)

func GetLivro(idLivro int) (*model.Livro, error) {
	var livro model.Livro
	resultado := db.DB.First(&livro, idLivro)

	if resultado.Error != nil {
		return nil, resultado.Error
	}
	return &livro, nil
}

func CriarLivro(livro *model.Livro) (uint, error) {
	resultado := db.DB.Create(&livro)

	if resultado.Error != nil {
		return 0, nil
	}
	return livro.ID, nil
}
