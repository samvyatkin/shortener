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
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	resAddr := fmt.Sprintf("http://%s", req.Host)
	if addr := h.config.ShortenerAddr; addr != nil {
		resAddr = *addr
	}

	UUID := h.uuidGenerator.Generate()
	shortURL := fmt.Sprintf("%s/%s", resAddr, UUID)

	url := string(body)
	url = strings.TrimSpace(url)

	h.storage.Set(UUID, url)

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)

	res.Write([]byte(shortURL))
}
