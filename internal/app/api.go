package api

import "net/http"

type Handler struct {
	Get    func(key string) ([]string, bool)
	Put    func(key, value string)
	Delete func(key string) bool
}

func New(*Handler) {
	panic("NYI")
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	panic("NYI")
}

func (h *Handler) PutHandler(w http.ResponseWriter, r *http.Request) {
	panic("NYI")
}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	panic("NYI")
}
