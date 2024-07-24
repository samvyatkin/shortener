package create

import (
	"fmt"
	"io"
	"net/http"
	"shortener/internal/app/storage"
	"shortener/internal/app/utils"
	"strconv"
	"strings"
)

type Handler struct {
	uuidGenerator utils.IdentifierGenerator
	storage       storage.Storage
}

func New(
	uuidGenerator utils.IdentifierGenerator,
	storage storage.Storage,
) *Handler {
	return &Handler{
		uuidGenerator: uuidGenerator,
		storage:       storage,
	}
}

func (h *Handler) Handle(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		h.getSourceURL(res, req)
	case http.MethodPost:
		h.createShortURL(res, req)
	}
}

func (h *Handler) createShortURL(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	UUID := h.uuidGenerator.Generate()
	shortURL := fmt.Sprintf("http://%s/%s", req.Host, UUID)

	url := string(body)
	url = strings.TrimSpace(url)

	h.storage.Set(UUID, url)

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)

	res.Write([]byte(shortURL))
}

func (h *Handler) getSourceURL(res http.ResponseWriter, req *http.Request) {
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
