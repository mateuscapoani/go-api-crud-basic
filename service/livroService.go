package service

import (
	"github.com/mateuscapoani/go-api-crud-basic/model"
	"github.com/mateuscapoani/go-api-crud-basic/repository"
)

func GetLivro(idLivro int) *model.Livro {
	return repository.GetLivro(idLivro)
}

func CriarLivro(livro *model.Livro) {
	repository.CriarLivro(livro)
}
