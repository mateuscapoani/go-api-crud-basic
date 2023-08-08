package request

type LivroRequest struct {
	Nome    string `json:"nome"`
	Autor   string `json:"autor"`
	Editora string `json:"editora"`
}
