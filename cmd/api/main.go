package main

import (
	"fmt"

	"github.com/mateuscapoani/go-api-crud-basic/config"
	"github.com/mateuscapoani/go-api-crud-basic/db"
)

func init() {
	config.CarregarVariaveisDeAmbiente()
	db.CriarConexao()
}

func main() {
	var result int
	db.DB.Raw("select 1+1").Scan(&result)
	fmt.Println(result)
}
