package model

import "gorm.io/gorm"

type Livro struct {
	gorm.Model
	Nome    string
	Autor   string
	Editora string
}
