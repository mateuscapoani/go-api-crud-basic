package main

import (
	"github.com/mateuscapoani/go-api-crud-basic/config"
	"github.com/mateuscapoani/go-api-crud-basic/db"
)

func init() {
	config.CarregarVariaveisDeAmbiente()
	db.CriarConexao()
}

func main() {
	config.IniciarRoteamento()
}
