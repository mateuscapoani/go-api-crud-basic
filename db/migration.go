package db

import "github.com/mateuscapoani/go-api-crud-basic/model"

func AutoMigration() {
	DB.AutoMigrate(&model.Livro{})
}
