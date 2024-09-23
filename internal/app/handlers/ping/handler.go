package ping

import (
	"net/http"
	"shortener/internal/app/config"
	"shortener/internal/app/storage"
)

type Handler struct {
	storage storage.Storage
	config  config.Configuration
}

func New(
	storage storage.Storage,
	config config.Configuration,
) *Handler {
	return &Handler{
		storage: storage,
		config:  config,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	err := h.storage.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer h.storage.Close()

	w.WriteHeader(http.StatusOK)
}
