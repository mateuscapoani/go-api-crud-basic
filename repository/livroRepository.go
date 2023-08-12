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
		return 0, resultado.Error
	}

	return livro.ID, nil
}

func EditarLivro(livro *model.Livro) (uint, error) {
	resultado := db.DB.Save(&livro)
	if resultado.Error != nil {
		return 0, resultado.Error
	}

	return livro.ID, nil
}

func DeletarLivro(idLivro int) (uint, error) {
	resultado := db.DB.Delete(&model.Livro{}, idLivro)
	if resultado.Error != nil {
		return 0, resultado.Error
	}

	return uint(idLivro), nil
}
