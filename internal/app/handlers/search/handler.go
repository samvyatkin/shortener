package search

import (
	"net/http"
	"shortener/internal/app/storage"
	"strconv"
	"strings"
)

type Handler struct {
	storage storage.Storage
}

func New(storage storage.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) Handle(res http.ResponseWriter, req *http.Request) {
	if u := req.URL; u != nil {
		params := strings.Split(u.Path, "/")
		for _, param := range params {
			url, err := h.storage.Get(param)
			if err == nil {
				res.Header().Set("Content-Type", "text/plain")
				res.Header().Set("Content-Length", strconv.Itoa(len(url)))
				res.Header().Set("Location", url)

				res.WriteHeader(http.StatusTemporaryRedirect)
				return
			}
		}
	}

	res.WriteHeader(http.StatusBadRequest)
}
