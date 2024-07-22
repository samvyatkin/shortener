package shortener_get

import "net/http"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(res http.ResponseWriter, req *http.Request) {}
