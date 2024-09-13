package shorten

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"shortener/internal/app/config"
	"shortener/internal/app/models"
	"shortener/internal/app/storage"
	"shortener/internal/app/utils"
)

type Handler struct {
	uuidGenerator *utils.UUIDGenerator
	storage       storage.Storage
	config        config.Configuration
}

func New(
	uuidGenerator *utils.UUIDGenerator,
	storage storage.Storage,
	config config.Configuration,
) *Handler {
	return &Handler{
		uuidGenerator: uuidGenerator,
		storage:       storage,
		config:        config,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var bodyData models.RequestData

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err = json.Unmarshal(buf.Bytes(), &bodyData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	UUID := h.uuidGenerator.Generate()
	shortURL := fmt.Sprintf("%s/%s", h.config.ShortenerAddr(), UUID)

	respData := models.ResponseData{ShortURL: shortURL}
	respBody, err := json.Marshal(respData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	d := models.ShortenData{
		ID:          UUID,
		ShortURL:    UUID,
		OriginalURL: bodyData.URL,
	}

	err = h.storage.Set(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(respBody)
}
