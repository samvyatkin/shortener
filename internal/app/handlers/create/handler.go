package create

import (
	"fmt"
	"io"
	"net/http"
	"shortener/internal/app/config"
	"shortener/internal/app/models"
	"shortener/internal/app/storage"
	"shortener/internal/app/utils"
	"strings"
)

type Handler struct {
	uuidGenerator   utils.IdentifierGenerator
	inMemoryStorage storage.Storage
	fileStorage     storage.Storage
	config          config.Configuration
}

func New(
	uuidGenerator utils.IdentifierGenerator,
	inMemoryStorage storage.Storage,
	fileStorage storage.Storage,
	config config.Configuration,
) *Handler {
	return &Handler{
		uuidGenerator:   uuidGenerator,
		inMemoryStorage: inMemoryStorage,
		fileStorage:     fileStorage,
		config:          config,
	}
}

func (h *Handler) Handle(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	defer req.Body.Close()

	UUID := h.uuidGenerator.Generate()
	shortURL := fmt.Sprintf("%s/%s", h.config.ShortenerAddr(), UUID)

	d := models.ShortenData{
		ID:          UUID,
		ShortURL:    UUID,
		OriginalURL: strings.TrimSpace(string(body)),
	}

	err = h.inMemoryStorage.Set(d)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.fileStorage.Set(d)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)

	res.Write([]byte(shortURL))
}
