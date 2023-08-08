package service

import (
	"github.com/mateuscapoani/go-api-crud-basic/dto/request"
	"github.com/mateuscapoani/go-api-crud-basic/dto/response"
	"github.com/mateuscapoani/go-api-crud-basic/model"
	"github.com/mateuscapoani/go-api-crud-basic/repository"
)

func GetLivro(idLivro int) (*response.LivroResponse, error) {
	livro, err := repository.GetLivro(idLivro)
	if err != nil {
		return nil, err
	}
	return &response.LivroResponse{ID: livro.ID, Nome: livro.Nome, Autor: livro.Autor, Editora: livro.Editora}, nil
}

func CriarLivro(LivroRequest *request.LivroRequest) (uint, error) {
	livro := model.Livro{Nome: LivroRequest.Nome, Autor: LivroRequest.Autor, Editora: LivroRequest.Editora}
	return repository.CriarLivro(&livro)
}
