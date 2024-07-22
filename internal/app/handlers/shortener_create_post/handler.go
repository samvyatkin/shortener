package shortener_create_post

import (
	"net/http"
	"strings"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getSourceURL(res, req)
	case http.MethodPost:
		createShortURL(res, req)
	}
}

func createShortURL(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
}

func getSourceURL(res http.ResponseWriter, req *http.Request) {
	params := strings.Split(req.URL.Path, "/")
	if ID := findID(params); ID != nil {
		res.Header().Set("Content-Type", "text/plain")
		res.Header().Set("Location", "https://google.com")

		res.WriteHeader(http.StatusTemporaryRedirect)
		return
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
