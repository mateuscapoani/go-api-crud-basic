package response

type LivroResponse struct {
	ID      uint   `json:"id"`
	Nome    string `json:"nome"`
	Autor   string `json:"autor"`
	Editora string `json:"editora"`
}
