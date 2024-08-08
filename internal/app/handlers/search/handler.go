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
	u := req.URL
	if u == nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := h.findURL(strings.Split(u.Path, "/"))
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.Header().Set("Content-Length", strconv.Itoa(len(url)))
	res.Header().Set("Location", url)

	res.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) findURL(params []string) (string, error) {
	for _, param := range params {
		url, err := h.storage.Get(param)
		if err != nil {
			continue
		}

		return url, nil
	}

	return "", storage.ErrStorageValueNotFound
}
