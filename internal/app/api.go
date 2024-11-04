package api

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Get    func(key string) ([]string, bool)
	Put    func(key, value string)
	Delete func(key string) bool
}

func New(*Handler) {
	panic("NYI")
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, exists := h.Get(key)
	if !exists {
		http.Error(w, "key not found", http.StatusNotFound)
		return
	}
	jsonData, err := json.Marshal(value)
	if err != nil {
		http.Error(w, "error encoding json", http.StatusInternalServerError)
		return
	}
	w.WriteHeader((http.StatusOK))
	w.Write(jsonData)
}

func (h *Handler) PutHandler(w http.ResponseWriter, r *http.Request) {
	panic("NYI")
}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	panic("NYI")
}
