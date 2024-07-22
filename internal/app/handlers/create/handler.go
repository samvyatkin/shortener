package create

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	cache map[string]string
}

func New() *Handler {
	return &Handler{
		cache: make(map[string]string),
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

	UUID := uuid.New().String()
	shortURL := fmt.Sprintf("http://%s/%s", req.Host, UUID)

	url := string(body)
	url = strings.TrimSpace(url)

	h.cache[UUID] = url

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)

	res.Write([]byte(shortURL))
}

func (h *Handler) getSourceURL(res http.ResponseWriter, req *http.Request) {
	if u := req.URL; u != nil {
		params := strings.Split(u.Path, "/")
		for _, param := range params {
			if cached, ok := h.cache[param]; ok {
				res.Header().Set("Content-Type", "text/plain")
				res.Header().Set("Content-Length", strconv.Itoa(len(cached)))
				res.Header().Set("Location", cached)

				res.WriteHeader(http.StatusTemporaryRedirect)
				return
			}
		}
	}

	res.WriteHeader(http.StatusBadRequest)
}
