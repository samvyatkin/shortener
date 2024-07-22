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
	fmt.Println(req.Method)

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

	fmt.Println(body)

	url := string(body)
	url = strings.TrimSpace(url)

	h.cache[url] = url

	fmt.Println(h.cache)

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
}

func (h *Handler) getSourceURL(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Println("Parse Body Parameters")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	params := strings.Split(req.URL.Path, "/")
	fmt.Println(params)
	if ID := findID(params); ID != nil {
		if cached, ok := h.cache[*ID]; ok {
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

func findID(params []string) *string {
	for _, param := range params {
		if param != "" {
			return &param
		}
	}

	return nil
}
