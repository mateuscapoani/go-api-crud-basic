package response

type erroResponse struct {
	Mensagem string `json:"mensagem"`
}

func BadRequestResponse(mensagem string) (int, erroResponse) {
	return erroBase(400, mensagem)
}

func NotFoundResponse(mensagem string) (int, erroResponse) {
	return erroBase(404, mensagem)
}

func erroBase(statusCode int, mensagem string) (int, erroResponse) {
	return statusCode, erroResponse{Mensagem: mensagem}
}
