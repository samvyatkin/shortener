package create

import (
	"fmt"
	"io"
	"net/http"
	"shortener/internal/app/config"
	"shortener/internal/app/storage"
	"shortener/internal/app/utils"
	"strings"
)

type Handler struct {
	uuidGenerator utils.IdentifierGenerator
	storage       storage.Storage
	config        config.Configuration
}

func New(
	uuidGenerator utils.IdentifierGenerator,
	storage storage.Storage,
	config config.Configuration,
) *Handler {
	return &Handler{
		uuidGenerator: uuidGenerator,
		storage:       storage,
		config:        config,
	}
}

func (h *Handler) Handle(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer req.Body.Close()

	UUID := h.uuidGenerator.Generate()
	shortURL := fmt.Sprintf("%s/%s", h.config.ShortenerAddr(), UUID)

	url := string(body)
	url = strings.TrimSpace(url)

	h.storage.Set(UUID, url)

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)

	res.Write([]byte(shortURL))
}
