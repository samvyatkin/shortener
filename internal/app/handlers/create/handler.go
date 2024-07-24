package create

import (
	"fmt"
	"io"
	"net/http"
	"shortener/internal/app/storage"
	"shortener/internal/app/utils"
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
	defer req.Body.Close()
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
