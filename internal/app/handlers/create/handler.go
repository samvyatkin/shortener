package create

import (
	"fmt"
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

	url := string(body)
	url = strings.TrimSpace(url)

	h.cache[url] = url

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
}

func (h *Handler) getSourceURL(res http.ResponseWriter, req *http.Request) {
	if u := req.URL; u != nil {
		if cached, ok := h.cache[u.String()]; ok {
			fmt.Println(cached)

			res.Header().Set("Content-Type", "text/plain")
			res.Header().Set("Content-Length", strconv.Itoa(len(cached)))
			res.Header().Set("Location", cached)

			res.WriteHeader(http.StatusTemporaryRedirect)
			return
		}
	}

	res.WriteHeader(http.StatusBadRequest)
}
