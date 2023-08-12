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

func CriarLivro(livroRequest *request.LivroRequest) (uint, error) {
	livro := model.Livro{Nome: livroRequest.Nome, Autor: livroRequest.Autor, Editora: livroRequest.Editora}
	return repository.CriarLivro(&livro)
}

func EditarLivro(idLivro int, livroRequest *request.LivroRequest) (uint, error) {
	livro, err := repository.GetLivro(idLivro)
	if err != nil {
		return 0, err
	}

	livro.Nome = livroRequest.Nome
	livro.Autor = livroRequest.Autor
	livro.Editora = livroRequest.Editora

	return repository.EditarLivro(livro)
}

func DeletarLivro(idLivro int) (uint, error) {
	_, err := repository.GetLivro(idLivro)
	if err != nil {
		return 0, err
	}

	return repository.DeletarLivro(idLivro)
}
